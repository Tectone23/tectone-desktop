import { FormEvent, useEffect, useRef, useState } from "react";
import {
  CreateProjectFromExampleProject,
  GetAllProjects,
  GetProjectByID,
} from "../../wailsjs/go/main/Project";
import { EventsOn } from "../../wailsjs/runtime/runtime";
import { main } from "../../wailsjs/go/models";
import Sidebar from "../components/Sidebar";
import Logo from "../assets/images/White_Logo.webp";
import { Label, Select } from "flowbite-react";
import ProjectSelect from "../components/ProjectSelect";
import { GetAllExampleProjectNames } from "../../wailsjs/go/main/SDK";
import { useProjects } from "../hooks/useProjects";
import Container from "./Container";
import Heading from "../components/Heading";

export default function Sandbox() {
  const { projects } = useProjects();
  const [selectedProject, setSelectedProject] = useState<string | null>(null);
  const [selectedExampleProject, setSelectedExampleProject] = useState<
    string | null
  >(null);
  const [project, setProject] = useState<main.Project | null>(null);
  const [stdout, setStdout] = useState<string[]>([]);
  const [examples, setExamples] = useState<string[]>([]);
  const stdoutRef = useRef<HTMLDivElement>(null);

  // EventsOn("sandbox-up-process", (data) => {
  //   setStdout((state) => {
  //     const newState = [data, ...state];
  //     return newState;
  //   });
  // });

  EventsOn("scaffold-example-stdout", (data: string) => {
    setStdout((state) => {
      const newState = [...state, data];
      return newState;
    });
  });

  useEffect(() => {
    if (!selectedProject || !project) return;

    async function getProjectById() {
      const p = await GetProjectByID(selectedProject!!);
      setProject(p);
    }

    async function loadSampleProjects() {
      const sp = await GetAllExampleProjectNames(project?.sdk_path!!);
      setExamples(sp);
    }

    getProjectById();
    loadSampleProjects();
  }, [selectedProject]);
  useEffect(() => {
    stdoutRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [stdout]);

  function handleSelectedOption(e: any) {
    const val = e.target.value;
    if (val == "") {
      setSelectedExampleProject(null);
      setSelectedProject(null);
      setProject(null);
      setExamples([]);
      return;
    }
    setSelectedProject(val);
    const p = projects.filter((p) => p.id == val)[0];
    setProject(p);
  }

  async function scaffoldExampleProject() {
    if (!selectedExampleProject) return;
    await CreateProjectFromExampleProject(selectedExampleProject);
  }

  return (
    <Container>
      <Heading
        title="Example Projects"
        description="
          Select an example project that is suitable for your project to get
          quickly started.
      "
      />
      <div className="flex flex-col w-full max-h-full overflow-y-scroll text-gray-300">
        <div>
          <label
            htmlFor="projects"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Select Projects
          </label>
          <select
            id="projects"
            className="block w-full p-2 mb-6 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            onChange={handleSelectedOption}
          >
            <option selected value={""}>
              Choose an existing Project
            </option>
            {projects?.map((proj, index) => {
              return (
                <option key={index} value={proj?.id}>
                  {proj?.name}
                </option>
              );
            })}
          </select>
        </div>
        {/* FIXME: (abdisamad) - replace the above select with the one below that is currently commented out */}
        {/* <div className="max-w-md">
          <ProjectSelect projects={projects} />
        </div> */}
        <div className="w-full flex space-x-4">
          <div className="space-y-2 px-4 py-2 flex flex-col ">
            {examples.map((example, index) => (
              <div
                className={`flex w-full space-x-2 ${
                  example == selectedExampleProject
                    ? "text-gray-200"
                    : " text-gray-600 hover:text-gray-200"
                }  text-sm`}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={1.5}
                  stroke="currentColor"
                  className="w-6 h-6 "
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44Z"
                  />
                </svg>

                <button
                  className={` ${
                    example == selectedExampleProject
                      ? "text-gray-100"
                      : "text-gray-400  hover:text-gray-100"
                  } `}
                  key={index}
                  onClick={() => setSelectedExampleProject(example)}
                >
                  {example}
                </button>

                {example == selectedExampleProject && (
                  <div className="flex items-center text-bold text-xs text-teal-500">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth={1.5}
                      stroke="currentColor"
                      className="w-4 h-4"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M6.75 15.75 3 12m0 0 3.75-3.75M3 12h18"
                      />
                    </svg>
                  </div>
                )}
              </div>
            ))}
          </div>
          {stdout.length > 0 && (
            <div className="flex flex-1 px-4 h-full w-full pt-2">
              <div className="flex h-full w-full bg-black/50 px-4 py-2 rounded-lg text-gray-600 ">
                {stdout.map((line, index) => {
                  return <pre key={index}>{line}</pre>;
                })}
                <div ref={stdoutRef} />
              </div>
            </div>
          )}
        </div>
        {selectedProject && (
          <div className="w-full flex justify-end items-center px-4 py-2">
            <button
              onClick={() => {
                scaffoldExampleProject();
              }}
              className={`${
                selectedExampleProject == "" || !selectedExampleProject
                  ? " bg-gray-600 cursor-not-allowed "
                  : "border border-gray-500 hover:bg-gray-700"
              } rounded-md  px-3 py-2 `}
            >
              Create Project
            </button>
          </div>
        )}
      </div>
    </Container>
  );
}
