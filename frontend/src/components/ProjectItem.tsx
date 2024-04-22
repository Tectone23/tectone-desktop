import { useEffect, useState } from "react";
import { main } from "../../wailsjs/go/models";
import {
  DockerComposeUp,
  DockerComposeDown,
  GetContainerStatus,
} from "../../wailsjs/go/main/Docker";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import {
  Cog6ToothIcon,
  EllipsisVerticalIcon,
  PlayIcon,
  StopIcon,
  TrashIcon,
} from "@heroicons/react/20/solid";
import { Navigate, useNavigate } from "react-router-dom";
import ContainerControls from "./ContainerControls";
import { DeleteProjectByID } from "../../wailsjs/go/main/Project";

interface ProjectItemProps {
  project: main.Project;
}
export default function ProjectItem({ project }: ProjectItemProps) {
  const [menuOpen, setMenuOpen] = useState<boolean>(false);
  // const [stdout, setStdout] = useState<string[]>([]);
  const [status, setStatus] = useState<string>("");
  const [isDeleteConfirmationOpen, setIsDeleteConfirmationOpen] =
    useState<boolean>(false);
  const navigate = useNavigate();

  // EventsOn(`sandbox-up-${project.id}`, (line: string) => {
  //   setStdout((state) => {
  //     const newState = [...state, line];
  //     return newState;
  //   });
  // });

  // EventsOn(`sandbox-down-${project.id}`, (line: string) => {
  //   setStdout((state) => {
  //     const newState = [...state, line];
  //     return newState;
  //   });
  // });

  useEffect(() => {
    async function getContainerStatus() {
      const statuses = await Promise.all(
        project.containers.map(async (container) => {
          const resp = await GetContainerStatus(container.id);
          console.log({ id: container.id, status: resp });
          return resp;
        })
      );

      const running = statuses.filter((v) => v == "running");
      const exited = statuses.filter((v) => v == "exited");
      if (running.length === 4) setStatus(() => "running");
      if (exited.length === 4) setStatus("exited");
    }

    const n = setInterval(async () => {
      await getContainerStatus();
    }, 2000);
    return () => {
      clearInterval(n);
    };
  }, []);

  async function updateProject() {
    return;
  }
  async function startSandbox() {
    // setStdout([]);
    setStatus("starting");
    await DockerComposeUp(project);
    return;
  }
  async function stopSandbox() {
    // setStdout([]);
    setStatus("stopping");
    await DockerComposeDown(project);
    return;
  }

  async function deleteProject() {
    await DeleteProjectByID(project.id);
  }

  function goToProjectView() {
    navigate(`/project-view`, {
      state: {
        projectId: project.id,
      },
    });
  }

  function goToProjectSettings() {
    navigate(`/project-settings`, {
      state: {
        projectId: project.id,
      },
    });
  }

  return (
    <>
      {isDeleteConfirmationOpen && (
        <div className="w-[200px] p-4 border-gray-200">
          are you sure you want to delete?
        </div>
      )}
      {/* <div className="flex w-full flex-col px-4 pb-8  rounded-sm border border-gray-600 shadow-sm">
        <div className="flex w-full justify-between border-b border-gray-700 mb-4 items-center py-2">
          <div className="font-semibold text-white">New Project</div>
          <div className="flex space-x-2">
            <button
              onClick={startSandbox}
              className="border-gray-600 border text-gray-200 hover:text-white py-2 px-3 text-xs  rounded-md"
            >
              Start sandbox
            </button>
            <button
              onClick={stopSandbox}
              className="border-gray-600 border text-gray-200 hover:text-white py-2 px-3 text-xs  rounded-md"
            >
              Stop Standbox
            </button>
            <button className="border-gray-600 border text-gray-200 hover:text-white py-2 px-3 text-xs  rounded-md">
              Update Project
            </button>
            <button onClick={() => setMenuOpen(!menuOpen)}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                stroke="currentColor"
                className="w-6 h-6"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z"
                />
              </svg>
            </button>
          </div>
        </div>
        <div className="flex flex-col space-y-8 w-full ">
          <div className="flex flex-col">
            <div className="text-gray-200 font-thin text-normal">
              Project name:
            </div>
            <span className="text-white font-thin text-normal">
              {project.name}
            </span>
          </div>
          <div className="flex flex-col">
            <div className="flex text-gray-200 font-thin text-normal">
              Project directory:
            </div>
            <span className="text-white font-thin text-normal">
              {project.path}
            </span>
          </div>
        </div>
        {stdout.length > 0 && (
          <div className="flex">
            <div className="h-[300px] w-full flex flex-col rounded-md bg-black overflow-y-scroll">
              {stdout.map((line) => (
                <pre className="text-gray-400">{line}</pre>
              ))}
            </div>
          </div>
        )}
      </div> */}

      <tr className="hover:cursor-pointer text-gray-500  px-4">
        <td className="whitespace-nowrap w-4 py-4 pl-4 pr-3 text-sm font-medium hover:text-gray-200 text-gray-500 sm:pl-0">
          {/* edit */}
          <Cog6ToothIcon
            className="w-4 h-4 "
            onClick={() => goToProjectSettings()}
          />
        </td>
        <td
          className="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-white sm:pl-0"
          onClick={() => {
            goToProjectView();
          }}
        >
          {project.name}
        </td>
        <td className="whitespace-nowrap text-ellipsis px-3 py-4 text-sm text-gray-300">
          {project.path}
        </td>
        <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-300">
          {status}
        </td>
        <td className="whitespace-nowrap px-3 py-4 text-sm text-gray-300">
          <ContainerControls
            onStart={startSandbox}
            onStop={stopSandbox}
            status={status}
          />
        </td>
        <td className="">
          <button>
            <EllipsisVerticalIcon className="text-gray-400 hover:text-gray-200 h-4 w-4" />
            <span className="sr-only">Option</span>
          </button>
          <button className=" text-red-800 hover:text-red-600">
            <TrashIcon className="h-4 w-4" onClick={() => deleteProject()} />
          </button>
        </td>
      </tr>
    </>
  );
}
