import ProjectItem from "./ProjectItem";
import { useProjects } from "../hooks/useProjects";
import { main } from "../../wailsjs/go/models";

interface ProjectsListContainerProps {
  projects: main.Project[];
}
export default function ProjectList({ projects }: ProjectsListContainerProps) {
  return (
    <>
      <div className="w-full py-8">
        {!projects || projects?.length == 0 ? (
          <div className="w-full h-[400px] flex justify-center items-center">
            <div className="border border-gray-700 mx-auto px-32 border-dashed text-gray-500 py-16 rounded-md">
              No projects currently
            </div>
          </div>
        ) : (
          // <div>
          //   <div>
          //     {projects.map((project) => (
          //       <ProjectItem project={project} />
          //     ))}
          //   </div>
          // </div>
          <table className="min-w-full divide-y divide-gray-700">
            <thead>
              <tr className="uppercase">
                <th
                  scope="col"
                  className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-white sm:pl-0"
                >
                  Edit
                </th>
                <th
                  scope="col"
                  className="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-white sm:pl-0"
                >
                  Name
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-white"
                >
                  Work dir
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-white"
                >
                  Status
                </th>
                <th
                  scope="col"
                  className="px-3 py-3.5 text-left text-sm font-semibold text-white"
                >
                  Actions
                </th>
                <th scope="col" className="relative py-3.5 pl-3 pr-4 sm:pr-0">
                  <span className="sr-only">Options</span>
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-800">
              {projects.map((project) => (
                <ProjectItem key={project.id} project={project} />
              ))}
            </tbody>
          </table>
        )}
      </div>
    </>
  );
}
