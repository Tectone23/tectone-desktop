import { useEffect, useState } from "react";
import { EventsOn } from "../../wailsjs/runtime/runtime";

export function useError() {
  const [error, setError] = useState<string>("");

  useEffect(() => {
    EventsOn("error", (msg: string) => {
      setError(msg);
    });
  }, [setError]);

  return {
    error,
    setError,
  };
}

export function useInfo() {
  const [info, setInfo] = useState<string>("");

  useEffect(() => {
    EventsOn("info", (msg: string) => {
      setInfo(msg);
    });
  }, [setInfo]);

  return {
    info,
    setInfo,
  };
}
