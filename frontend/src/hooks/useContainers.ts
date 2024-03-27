import { useCallback, useEffect, useRef, useState } from "react";
import { useProject } from "./useProjects";
import {
  DockerComposeDown,
  DockerComposeUp,
  GetContainerStatus,
  StreamComposeLogs,
  StreamContainerLogs,
} from "../../wailsjs/go/main/Docker";
import { EventsOn } from "../../wailsjs/runtime/runtime";

export function useContainerStatus(projectId: string): {
  status: string;
  logs: string[];
  logsRef: React.RefObject<HTMLDivElement>;
  getContainerStatus: () => void;
  streamContainerLogs: () => void;
  startSandbox: () => void;
  stopSandbox: () => void;
} {
  const { project } = useProject(projectId);
  const [status, setStatus] = useState<string>("");
  const [logs, setLogs] = useState<string[]>([]);
  const logsRef = useRef<HTMLDivElement>(null);

  const getContainerStatus = useCallback(async () => {
    if (!project) return;
    const statuses = await Promise.all(
      project.containers.map(async (container) => {
        const resp = await GetContainerStatus(container.id);
        // console.log({ id: container.id, status: resp });
        return resp;
      })
    );

    const running = statuses.filter((v) => v == "running");
    const exited = statuses.filter((v) => v == "exited");
    if (running.length === 4) setStatus(() => "running");
    if (exited.length === 4) setStatus("exited");
  }, [project]);

  const streamContainerLogs = useCallback(async () => {
    console.log("get container logs");
    if (!project) return;
    console.log("getting container logs");
    setLogs([]);
    await StreamComposeLogs(project);
  }, [project]);

  const startSandbox = useCallback(async () => {
    setStatus("starting");
    await DockerComposeUp(project!!);
    return;
  }, []);

  const stopSandbox = useCallback(async () => {
    setStatus("stopping");
    await DockerComposeDown(project!!);
    return;
  }, []);

  useEffect(() => {
    if (!project) return;
    EventsOn("project-container-logs", (data: string) => {
      setLogs((state) => {
        const newState = [...state, data];
        return newState;
      });
    });

    getContainerStatus();
  }, [project, projectId]);

  return {
    status,
    logs,
    logsRef,
    getContainerStatus,
    streamContainerLogs,
    startSandbox,
    stopSandbox,
  };
}
