import Sidebar from "../components/Sidebar";
import Logo from "../assets/images/White_Logo.webp";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import DispenserForm from "../components/DispenserForm";
import NetworkConfig from "../components/NetworkConfig";
import { useProject } from "../hooks/useProjects";
import { useNetworkConfig, useNewProjectConfig } from "../hooks/useConfig";
// import { useApplictionStore } from "../store/store";

const tabs = ["project", "testnet", "mainnet", "devnet", "dispenser"];

export default function ProjectSettings() {
  const location = useLocation();
  const projectId = location?.state?.projectId;

  const { projectName, setProjectName, saveProjectConfig } =
    useNewProjectConfig(projectId);
  const { project, saveTestNetConfig, saveMainNetConfig, saveDevNetConfig } =
    useNetworkConfig(projectId);

  const [selectedTab, setSelectedTab] = useState<string>(tabs[0]);

  return (
    <>
      <div className="w-[300px] border-r border-gray-800 pt-8  bg-gray-900 h-full px-6">
        <div className="flex justify-between items-center  pb-12 border-b border-gray-700 mb-12">
          <img src={Logo} alt="" />
        </div>
        <Sidebar />
      </div>
      {project && (
        <div className="flex flex-col  w-full max-h-full overflow-y-scroll text-white px-4 mt-24">
          <div className="flex text-md font-semibold mb-6">
            Project Configuration
          </div>
          <div className="block mb-10">
            <p className="text-gray-300 text-sm max-w-2xl">
              View, edit, save, and add values to the genesis file for{" "}
              <strong className="text-white">TestNet</strong>,{" "}
              <strong className="text-white">MainNet</strong>, and{" "}
              <strong className="text-white">DevNet</strong>.
            </p>
          </div>
          <div className="flex flex-col space-y-4 text-sm">
            <div className="capitalize">
              Project:{" "}
              <span className="text-gray-400 text-sm">{project?.name}</span>
            </div>
            <div>
              Directory:{" "}
              <span className="text-gray-400 text-sm">{project?.path}</span>
            </div>
          </div>
          <div className="mt-12 w-full max-h-full">
            <div className="mb-4 border-b border-gray-200 dark:border-gray-700">
              <ul
                className="flex flex-wrap -mb-px text-sm font-medium text-center"
                id="default-tab"
                data-tabs-toggle="#default-tab-content"
                role="tablist"
              >
                <li className="me-2" role="presentation">
                  <button
                    className={`inline-block p-4 ${
                      selectedTab == "project" ? "border-b-2" : "text-gray-300"
                    }  rounded-t-lg`}
                    id="project-tab"
                    data-tabs-target="#project"
                    type="button"
                    role="tab"
                    aria-controls="project"
                    aria-selected={selectedTab == "project"}
                    onClick={() => setSelectedTab("project")}
                  >
                    Project
                  </button>
                </li>
                <li className="me-2" role="presentation">
                  <button
                    className={`inline-block p-4 ${
                      selectedTab == "testnet" ? "border-b-2" : "text-gray-300"
                    }  rounded-t-lg`}
                    id="testnet-tab"
                    data-tabs-target="#testnet"
                    type="button"
                    role="tab"
                    aria-controls="testnet"
                    aria-selected={selectedTab == "testnet"}
                    onClick={() => setSelectedTab("testnet")}
                  >
                    TestNet
                  </button>
                </li>
                <li className="me-2" role="presentation">
                  <button
                    className={`inline-block p-4 ${
                      selectedTab == "mainnet" ? "border-b-2" : "text-gray-300"
                    } rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300`}
                    id="mainnet-tab"
                    data-tabs-target="#mainnet"
                    type="button"
                    role="tab"
                    aria-controls="mainnet"
                    aria-selected={selectedTab == "mainnet"}
                    onClick={() => setSelectedTab("mainnet")}
                  >
                    MainNet
                  </button>
                </li>
                <li className="me-2" role="presentation">
                  <button
                    className={`inline-block p-4 ${
                      selectedTab == "devnet" ? "border-b-2" : "text-gray-300"
                    } rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300`}
                    id="devnet-tab"
                    data-tabs-target="#devnet"
                    type="button"
                    role="tab"
                    aria-controls="devnet"
                    aria-selected={selectedTab == "devnet"}
                    onClick={() => setSelectedTab("devnet")}
                  >
                    DevNet
                  </button>
                </li>
                <li className="me-2" role="presentation">
                  <button
                    className={`inline-block p-4 ${
                      selectedTab == "dispenser"
                        ? "border-b-2"
                        : "text-gray-300"
                    } rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300`}
                    id="dispenser-tab"
                    data-tabs-target="#dispenser"
                    type="button"
                    role="tab"
                    aria-controls="dispenser"
                    aria-selected={selectedTab == "dispenser"}
                    onClick={() => setSelectedTab("dispenser")}
                  >
                    Dispenser
                  </button>
                </li>
              </ul>
            </div>
            {project && (
              <div id="default-tab-content" className=" overflow-y-scroll">
                <div
                  className={`${
                    selectedTab == "project" ? "" : "hidden"
                  } p-4 rounded-lg bg-gray-50 dark:bg-gray-800`}
                  id="project"
                  role="tabpanel"
                  aria-labelledby="project-tab"
                >
                  <div>
                    <label
                      htmlFor="projectName"
                      className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                    >
                      Project Name
                    </label>
                    <input
                      type="text"
                      id="projectName"
                      className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                      value={projectName}
                      onChange={(e) => setProjectName(e.target.value)}
                    />
                  </div>
                  <div className="flex w-full justify-end pt-4 text-md">
                    <button
                      className="px-3 py-2 rounded border border-gray-700 text-gray-300 hover:bg-gray-700 hover:border-gray-600"
                      onClick={() => {
                        const p = project;
                        p.name = projectName;
                        saveProjectConfig(p);
                      }}
                    >
                      Save
                    </button>
                  </div>
                </div>
                <div
                  className={`${
                    selectedTab == "testnet" ? "" : "hidden"
                  } p-4 rounded-lg bg-gray-50 dark:bg-gray-800`}
                  id="testnet"
                  role="tabpanel"
                  aria-labelledby="testnet-tab"
                >
                  <NetworkConfig
                    key={project?.testnet?.genesis?.network}
                    genesis={project?.testnet?.genesis}
                    onSave={saveTestNetConfig}
                  />
                </div>
                <div
                  className={`${
                    selectedTab == "mainnet" ? "" : "hidden"
                  } p-4 rounded-lg bg-gray-50 dark:bg-gray-800`}
                  id="mainnet"
                  role="tabpanel"
                  aria-labelledby="mainnet-tab"
                >
                  <NetworkConfig
                    key={project.mainnet.genesis.network}
                    genesis={project.mainnet.genesis}
                    onSave={saveMainNetConfig}
                  />
                </div>
                <div
                  className={`${
                    selectedTab == "devnet" ? "" : "hidden"
                  } p-4 rounded-lg bg-gray-50 dark:bg-gray-800`}
                  id="devnet"
                  role="tabpanel"
                  aria-labelledby="devnet-tab"
                >
                  <NetworkConfig
                    key={project.devnet.genesis.network}
                    genesis={project.devnet.genesis}
                    onSave={saveDevNetConfig}
                  />
                </div>
                <div
                  className={`${
                    selectedTab == "dispenser" ? "" : "hidden"
                  } p-4 rounded-lg bg-gray-50 dark:bg-gray-800`}
                  id="dispenser"
                  role="tabpanel"
                  aria-labelledby="dispenser-tab"
                >
                  <DispenserForm />
                </div>
              </div>
            )}
          </div>
        </div>
      )}
    </>
  );
}
