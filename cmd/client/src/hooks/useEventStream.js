import { useEffect } from "react";
import { useAuth } from "../auth/useAuth";
import { API_HOST } from "../config/dev";
import { queryClient } from "../query/queryClient";

export function useEventStream(){
    const { user } = useAuth();
    

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

    eventSource.onmessage = (message) => {
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
          console.log("beep")
          queryClient.invalidateQueries({
            queryKey: ["choreTemplates", user?.id],
          });
          break;
      }

      console.log(event);
    };

    eventSource.onerror = (error) => {
        console.error("EventSource error:", error);
    }

    return () => {
      eventSource.close();
    };
  }, [user, queryClient]);}