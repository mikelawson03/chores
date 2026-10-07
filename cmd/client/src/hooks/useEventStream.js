import { useEffect, useRef } from "react";
import { useAuth } from "../auth/useAuth";
import { API_HOST } from "../config/dev";
import { queryClient } from "../query/queryClient";

export function useEventStream(){
    const { user, refreshUser } = useAuth();
    const assignmentInvalidationTimer = useRef(null);
    const choreTemplateInvalidationTimer = useRef(null);

    useEffect(() => {
    if (!user) {
      return;
    }

    const eventSource = new EventSource(
      `${API_HOST}/events`,
      { withCredentials: true }
    );

    function scheduleAssignmentInvalidation() {
      if (assignmentInvalidationTimer.current == null) {
        assignmentInvalidationTimer.current = setTimeout(() => {
          queryClient.invalidateQueries({
            queryKey: ["assignments", user?.user.id]
          });

          assignmentInvalidationTimer.current = null;
        }, 50);
      }
    }

    function scheduleChoreTemplateInvalidation() {
      if (choreTemplateInvalidationTimer.current == null) {
        choreTemplateInvalidationTimer.current = setTimeout(() => {
          queryClient.invalidateQueries({
            queryKey: ["choreTemplates", user?.householdId]
          });

          choreTemplateInvalidationTimer.current = null;
        }, 50);
      }
    }

    eventSource.onopen = () => {
        console.log("EventSource opened")
    }

    eventSource.onmessage = async (message) => {
      const event = JSON.parse(message.data);
      switch (event.type) {
        case "assignment.edited":
        case "assignment.created":
        case "assignment.completed":
        case "assignment.rescheduled":
          scheduleAssignmentInvalidation();
          break;
        case "choreTemplate.edited":
        case "choreTemplate.created":
        case "choreTemplate.deleted":
          scheduleChoreTemplateInvalidation();
          break;
        case "householdUser.edited":
        case "householdUser.added":
        case "householdUser.deactivated":
        case "householdUser.activated":
        case "householdUser.removed":
          queryClient.invalidateQueries({
            queryKey: ["activeHouseholdUsers", user?.householdId],
          });

          if (event.payload.user.id === user?.user.id) {
            await refreshUser();
          }
          break;
        case "user.edited":
          queryClient.invalidateQueries({
            queryKey: ["activeHouseholdUsers", user?.householdId],
          });

          if (event.payload.id === user?.user.id) {
            await refreshUser();
          }
          break;
      }

    };

    eventSource.onerror = (error) => {
        console.error("EventSource error:", error);
    }

    return () => {
      eventSource.close();

      clearTimeout(assignmentInvalidationTimer.current);
      clearTimeout(choreTemplateInvalidationTimer.current);

      assignmentInvalidationTimer.current = null;
      choreTemplateInvalidationTimer.current = null;
    };
  }, [user, queryClient]);}