import { useEffect } from "react";
import { useAuth } from "../auth/useAuth";
import { API_HOST } from "../config/dev";
import { queryClient } from "../query/queryClient";

export function useEventStream(){
    const { user, refreshUser } = useAuth();
    

    useEffect(() => {
    if (!user) {
      return;
    }

    const eventSource = new EventSource(
      `${API_HOST}/events`,
      { withCredentials: true }
    );

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
          queryClient.invalidateQueries({
            queryKey: ["assignments", user?.id],
          });
          break;
        case "choreTemplate.edited":
        case "choreTemplate.created":
        case "choreTemplate.deleted":
          queryClient.invalidateQueries({
            queryKey: ["choreTemplates", user?.id],
          });
          break;
        case "householdUser.edited":
        case "householdUser.added":
        case "householdUser.deactivated":
        case "householdUser.activated":
        case "householdUser.removed":
          queryClient.invalidateQueries({
            queryKey: ["householdUsers", user?.householdId],
          });

          if (event.payload.user.id === user?.user.id) {
            await refreshUser();
          }
          break;
        case "user.edited":
          queryClient.invalidateQueries({
            queryKey: ["householdUsers", user?.householdId],
          });

          if (event.payload.id === user?.user.id) {
            await refreshUser();
          }
          break;
      }

      // console.log(event);
    };

    eventSource.onerror = (error) => {
        console.error("EventSource error:", error);
    }

    return () => {
      eventSource.close();
    };
  }, [user, queryClient]);}