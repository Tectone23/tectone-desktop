// import { create } from "zustand";
// import { model } from "../../wailsjs/go/models";
// import {
//   GetAllProjects,
//   GetProjectByID,
//   SaveTestNetGenesisConfig,
// } from "../../wailsjs/go/main/Project";
// import { SubmitFeedback } from "../../wailsjs/go/main/App";

// interface ApplicationStore {
//   version: string;
//   commitHash: string;
//   buildDate: string;
//   projects: model.Project[];
//   project: model.Project | null;
//   getProjects: () => void;
//   getProject: (projectId: string) => void;
//   saveTestNet: (project: model.Project, genesis: model.Genesis) => void;
//   updateProject: (project: model.Project) => void;
//   submitFeedback: (f: model.Feedback) => void;
// }

// export const useApplictionStore = create<ApplicationStore>()((set) => ({
//   version: "",
//   commitHash: "",
//   buildDate: "",
//   projects: [],
//   getProjects: async () => {
//     console.log("all projects executed");
//     const projects = await GetAllProjects();

//     set(() => {
//       return { projects: projects };
//     });
//   },
//   project: null,
//   getProject: async (projectId: string) => {
//     const p = await GetProjectByID(projectId);
//     set(() => ({
//       project: { ...p, convertValues: p.convertValues },
//     }));
//   },
//   saveTestNet: async (project: model.Project, genesis: model.Genesis) => {
//     const g = await SaveTestNetGenesisConfig(project, genesis);
//     set((state) => {
//       //   const p = state.project ?? new model.Project();
//       if (!state.project) {
//         return { project: state.project };
//       } else {
//         const p = Object.assign(
//           {},
//           {
//             ...state.project,
//             convertValues: project.convertValues,
//           }
//         );
//         p.testnet.genesis = g;
//         return {
//           project: p,
//         };
//       }
//     });
//   },
//   updateProject: async (p: model.Project) => {},
//   submitFeedback: async (f: model.Feedback) => {
//     const ok = await SubmitFeedback(f);
//   },
// }));

// interface ProjectConfigStore {
//   project: model.Project | null;
//   testnet: model.TestNet | null;
//   mainnet: model.MainNet | null;
//   devnet: model.DevNet | null;
//   getProject: (projectId: string) => void;
//   saveTestNet: (project: model.Project, genesis: model.Genesis) => void;
//   saveMainNet: (genesis: model.Genesis) => void;
//   saveDevNet: (genesis: model.Genesis) => void;
//   saveDispenser: (dispenser: any) => void;
// }

// const useProjectConfigStore = create<ProjectConfigStore>()((set) => ({
//   project: null,
//   testnet: null,
//   mainnet: null,
//   devnet: null,
//   getProject: async (projectId: string) => {
//     const p = await GetProjectByID(projectId);
//     set(() => {
//       return {
//         project: p,
//       };
//     });
//   },
//   saveTestNet: async (project: model.Project, genesis: model.Genesis) => {
//     const g = await SaveTestNetGenesisConfig(project, genesis);
//     set((state) => {
//       if (!state.project) {
//         return {
//           project: null,
//         };
//       }
//       return {
//         project: {
//           ...state.project,
//           convertValues: state.project.convertValues,
//           testnet: {
//             ...state.project.testnet,
//             convertValues: state.project.testnet.convertValues,
//             genesis: { ...g, convertValues: g.convertValues },
//           },
//         },
//       };
//     });
//   },
//   saveMainNet: () => {},
//   saveDevNet: () => {},
//   saveDispenser: () => {},
// }));

export const a = {};
