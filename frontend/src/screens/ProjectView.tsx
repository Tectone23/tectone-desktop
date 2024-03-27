import Sidebar from "../components/Sidebar";
import Logo from "../assets/images/White_Logo.webp";
import { useLocation, useNavigate } from "react-router-dom";
import TerminalView from "../components/TerminalView";
import { useEffect, useRef, useState } from "react";
import { GetProjectByID } from "../../wailsjs/go/main/Project";
import { main } from "../../wailsjs/go/models";
import ContainerControls from "../components/ContainerControls";
import {
  DockerComposeDown,
  DockerComposeUp,
  GetContainerStatus,
  StreamContainerLogs,
} from "../../wailsjs/go/main/Docker";
import { Cog6ToothIcon } from "@heroicons/react/20/solid";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { useProject } from "../hooks/useProjects";
import { useContainerStatus } from "../hooks/useContainers";

export default function ProjectView() {
  const location = useLocation();
  const projectId = location?.state?.projectId;
  const navigate = useNavigate();

  const { project } = useProject(projectId);
  const {
    status,
    getContainerStatus,
    logs,
    logsRef,
    startSandbox,
    stopSandbox,
    streamContainerLogs,
  } = useContainerStatus(projectId);

  const [logsSelected, setLogsSelected] = useState<boolean>(true);

  useEffect(() => {
    const n = setInterval(async () => {
      await getContainerStatus();
    }, 2000);
    return () => {
      clearInterval(n);
    };
  }, [project]);

  useEffect(() => {
    streamContainerLogs();
  }, [status]);

  useEffect(() => {
    logsRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [logs]);

  function selectLogs() {
    setLogsSelected(true);
  }

  function goToProjectSettings() {
    navigate(`/project-settings`, {
      state: {
        projectId: project?.id,
      },
    });
  }

  return (
    <>
      <div className="w-[300px] max-h-full  border-r border-gray-800 pt-8  bg-gray-900 h-full px-6">
        <div className="flex justify-between items-center  pb-12 border-b border-gray-700 mb-12">
          <img src={Logo} alt="" />
        </div>
        <Sidebar />
      </div>
      <div className="flex flex-col pt-20 w-full max-h-full overflow-y-hidden text-white px-4">
        <div className="block mt-4 mb-10">
          <p className="text-gray-300 text-sm max-w-2xl">View Logs and </p>
        </div>
        <div className="flex space-x-4">
          <div className="flex flex-col space-y-4">
            <div className="capitalize">
              Project:{" "}
              <span className="text-gray-400 text-sm">{project?.name}</span>
            </div>
            <div>
              Directory:{" "}
              <span className="text-gray-400 text-sm">{project?.path}</span>
            </div>
          </div>
        </div>
        <div className="flex-1 mt-12 w-full">
          <div className="flex justify-between mb-4 border-b items-center border-gray-200 dark:border-gray-700">
            <ul
              className="flex flex-wrap -mb-px text-sm font-medium text-center"
              id="default-tab"
              data-tabs-toggle="#default-tab-content"
              role="tablist"
            >
              <li className="me-2" role="presentation">
                <button
                  className={`inline-block p-4 ${
                    logsSelected ? "border-b-2" : "text-gray-300"
                  }  rounded-t-lg`}
                  id="logs-tab"
                  data-tabs-target="#logs"
                  type="button"
                  role="tab"
                  aria-controls="logs"
                  aria-selected={logsSelected}
                  onClick={() => selectLogs()}
                >
                  Logs
                </button>
              </li>
            </ul>
            <div className="flex justify-end space-x-4 items-end w-full h-full">
              <div className="flex py-5 text-sm text-gray-300 hover:text-white">
                <button
                  className="flex space-x-2"
                  onClick={() => goToProjectSettings()}
                >
                  <Cog6ToothIcon className="w-5 h-5 text-gray-200" />
                  <span>Project Configuration</span>
                </button>
              </div>
              <div className="py-2">
                <div className="border border-gray-700 rounded-md px-4 py-2">
                  <ContainerControls
                    onStart={startSandbox}
                    onStop={stopSandbox}
                    status={status}
                  />
                </div>
              </div>
            </div>
          </div>
          <div id="default-tab-content" className=" flex flex-col ">
            {/* <div className="mb-4 text-sm">Container logs streaming ...</div> */}
            <div
              className={`${
                logsSelected ? "" : "hidden"
              }  flex flex-1 w-full  p-4 rounded-lg bg-gray-50 dark:bg-gray-800`}
              id="logs"
              role="tabpanel"
              aria-labelledby="logs-tab"
            >
              {/* <TerminalView logs={logs} /> */}
              <div className="w-full  overflow-x-none flex flex-col overflow-y-scroll text-wrap h-[400px] max-h-[400px] text-xs bg-black/80 text-gray-500 font-mono rounded-md px-4 py-8">
                {logs.map((log, idx) => (
                  <pre className="text-wrap" key={idx}>
                    {log}
                  </pre>
                ))}
                <div ref={logsRef} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
