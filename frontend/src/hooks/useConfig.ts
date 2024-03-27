import { useCallback, useEffect, useState } from "react";
import { model } from "../../wailsjs/go/models";
import { useProject } from "./useProjects";
import {
  SaveDevNetGenesisConfig,
  SaveMainNetGenesisConfig,
  SaveProjectConfig,
  SaveTestNetGenesisConfig,
} from "../../wailsjs/go/main/Project";

export function useNetworkConfig(projectId: string): {
  project: model.Project | null;
  saveTestNetConfig: (genesis: model.Genesis) => void;
  saveMainNetConfig: (genesis: model.Genesis) => void;
  saveDevNetConfig: (genesis: model.Genesis) => void;
} {
  const { project, setProject } = useProject(projectId);

  useEffect(() => {
    console.log("new project loaded");
  });

  const saveTestNetConfig = useCallback(
    async (genesis: model.Genesis) => {
      if (!project) return;
      const g = await SaveTestNetGenesisConfig(project, genesis);
      const p = new model.Project(project);
      if (p.testnet.genesis) {
        console.log("setting project");
        p.testnet.genesis = new model.Genesis(g);
        setProject(p);
      }
    },
    [project]
  );

  const saveDevNetConfig = useCallback(
    async (genesis: model.Genesis) => {
      if (!project) return;
      const g = await SaveDevNetGenesisConfig(project, genesis);
      const p = new model.Project(project);
      if (p.devnet.genesis) {
        p.devnet.genesis = new model.Genesis(g);
        setProject(p);
      }
    },
    [project]
  );

  const saveMainNetConfig = useCallback(
    async (genesis: model.Genesis) => {
      if (!project) return;
      const g = await SaveMainNetGenesisConfig(project, genesis);
      const p = new model.Project(project);
      if (p.mainnet.genesis) {
        p.mainnet.genesis = new model.Genesis(g);
        setProject(p);
      }
    },
    [project]
  );

  return {
    project,
    saveTestNetConfig,
    saveDevNetConfig,
    saveMainNetConfig,
  };
}

export function useNewProjectConfig(projectId: string): {
  saveProjectConfig: (p: model.Project) => void;
  projectName: string;
  setProjectName: React.Dispatch<React.SetStateAction<string>>;
} {
  const [projectName, setProjectName] = useState<string>("");
  const { project } = useProject(projectId);

  const saveProjectConfig = useCallback(async (p: model.Project) => {
    await SaveProjectConfig(p);
  }, []);

  useEffect(() => {
    if (!project) return;
    setProjectName(project.name);
  }, [project]);

  return {
    saveProjectConfig,
    projectName,
    setProjectName,
  };
}
