import { useCallback, useEffect, useRef, useState } from "react";
import { model } from "../../wailsjs/go/models";
import {
  GetAllProjects,
  GetProjectByID,
  New,
  SetupProject,
} from "../../wailsjs/go/main/Project";
import { OpenDir } from "../../wailsjs/go/main/App";
import { EventsOn } from "../../wailsjs/runtime/runtime";

export function useProject(id: string): {
  project: model.Project | null;
  setProject: React.Dispatch<React.SetStateAction<model.Project | null>>;
} {
  const [project, setProject] = useState<model.Project | null>(null);

  useEffect(() => {
    async function getProject() {
      const p = await GetProjectByID(id);
      setProject(p);
    }
    getProject();
    console.log("running useProject");
  }, []);

  return {
    project,
    setProject,
  };
}

export function useCreateProject(projectDir: string): {
  projectName: string;
  setProjectName: React.Dispatch<React.SetStateAction<string>>;
  agreement: boolean;
  setAgreement: React.Dispatch<React.SetStateAction<boolean>>;
  logs: string[];
  logsRef: React.RefObject<HTMLDivElement>;
  inProgress: boolean;
  createProject: () => void;
} {
  const [projectName, setProjectName] = useState<string>("");
  const [agreement, setAgreement] = useState<boolean>(false);
  const [inProgress, setInProgress] = useState<boolean>(false);
  const [logs, setLogs] = useState<string[]>([]);
  const logsRef = useRef<HTMLDivElement>(null);

  const createProject = useCallback(async () => {
    if (projectDir == "") return;
    if (!agreement) {
      console.log("agreement is not checked!");
      return;
    }
    setInProgress(true);
    const path = await New(projectName, projectDir);
    const success = await SetupProject(projectName, path);
    setInProgress(false);
  }, [projectDir, agreement]);

  useEffect(() => {
    EventsOn("new-project-stdout", (data: string) => {
      setLogs((state) => {
        const newState = [...state, data];
        return newState;
      });
    });
  }, []);

  return {
    projectName,
    setProjectName,
    agreement,
    setAgreement,
    logs,
    logsRef,
    inProgress,
    createProject,
  };
}

export function useProjects(): { projects: model.Project[] } {
  const [projects, setProjects] = useState<model.Project[]>([]);

  const getProjects = useCallback(async () => {
    const ps = await GetAllProjects();
    setProjects(ps);
  }, []);

  useEffect(() => {
    getProjects();
  }, []);

  return {
    projects,
  };
}
