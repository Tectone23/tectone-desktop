import { NavLink } from "react-router-dom";
import { BrowserOpenURL } from "../../wailsjs/runtime/runtime";
import ProjectList from "../components/ProjectsListContainer";
import { useProjects } from "../hooks/useProjects";
import Container from "./Container";
import Heading from "../components/Heading";

export default function Home() {
  const { projects } = useProjects();
  // const projects = useApplictionStore((state) => state.projects);
  return (
    <Container>
      <div className="flex flex-col w-full">
        <div className="mt-4 mb-6  text-sm text-purple-400 underline flex space-x-4">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="w-6 h-6"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M17.25 8.25 21 12m0 0-3.75 3.75M21 12H3"
            />
          </svg>

          <button
            onClick={() => {
              BrowserOpenURL("https://developer.algorand.org/docs/");
            }}
          >
            Tectone Developer Documentation
          </button>
        </div>
        <div className="w-full flex justify-between">
          <Heading
            title="Projects"
            description="A list of all projects in your development environment."
          />
          <div className="sm:flex sm:items-center">
            <div className="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
              <NavLink
                to="/newProject"
                type="button"
                className="block rounded-md bg-gray-800 border border-gray-300 px-3 py-2 text-center text-sm font-semibold text-white hover:bg-gray-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
              >
                Add new project
              </NavLink>
            </div>
          </div>
        </div>

        <div className="mt-8 flow-root">
          <ProjectList projects={projects} />
        </div>
      </div>
    </Container>
  );
}
