import { useEffect } from "react";
import Spinner from "../components/Spinner";
import Sidebar from "../components/Sidebar";
import Logo from "../assets/images/White_Logo.webp";
import { useNavigate } from "react-router-dom";
import { useCreateProject } from "../hooks/useProjects";
import { useDirDialog } from "../hooks/useFs";

export default function NewProject() {
  const { projectDir, getProjectDir } = useDirDialog();
  const {
    projectName,
    setProjectName,
    agreement,
    setAgreement,
    logs,
    logsRef,
    inProgress,
    createProject,
  } = useCreateProject(projectDir);
  const navigate = useNavigate();

  useEffect(() => {
    logsRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [logs]);

  return (
    <>
      <div className="w-[300px] border-r border-gray-800 pt-8  bg-gray-900 h-full px-6">
        <div className="flex justify-between items-center  pb-12 border-b border-gray-700 mb-12">
          <img src={Logo} alt="" />
        </div>
        <Sidebar />
      </div>
      <div className="flex w-full  text-gray-200 px-4">
        <div className="flex flex-col  w-full mt-8 ">
          <div className="w-full flex flex-col pt-8 pb-12 border-b border-gray-700">
            <div className="flex text-md font-semibold mb-6">
              Project Configuration
            </div>
            <p className="flex justify-start text-gray-300 text-sm w-3/4">
              To get started with creating a new project give your project a
              name, and select a directory where your new project will be
              created.
            </p>
          </div>

          <div className="flex flex-col space-y-8 my-10">
            <div className="flex space-x-8 items-center w-full">
              <label
                htmlFor="project"
                className="block text-sm font-medium leading-6 text-gray-100"
              >
                Project Name
              </label>
              <div className="">
                <input
                  type="text"
                  name="project"
                  id="project"
                  className="block w-full placeholder:text-xs rounded-md border-0 pl-4 py-1.5 text-gray-100 bg-gray-800 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 placeholder:font-light focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  placeholder="New Project Name"
                  value={projectName}
                  onChange={(e) => setProjectName(e.target.value)}
                />
              </div>
            </div>
            <div className="flex space-x-8 items-center w-full">
              <label className="block text-sm font-medium leading-6 text-gray-100">
                Project Directory
              </label>
              <div className="flex space-x-12 items-center">
                <div
                  id="project-dir"
                  className="block min-w-48 border-b border-gray-500 pb-2 text-white"
                >
                  {projectDir}
                </div>
                <button
                  className="flex px-3 py-2 rounded-md text-gray-100 border border-gray-600 hover:bg-gray-700"
                  onClick={getProjectDir}
                >
                  Select Project Directory
                </button>
              </div>
            </div>
            <div className="flex space-x-8 items-center w-full">
              <div className="relative flex items-start">
                <div className="flex h-6 items-center">
                  <input
                    id="agreement"
                    aria-describedby="agreement-description"
                    type="checkbox"
                    className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
                    checked={agreement}
                    onChange={(e) => setAgreement(e.target.checked)}
                  />
                </div>
                <div className="ml-3 text-sm leading-6">
                  <label
                    htmlFor="agreement"
                    className="font-medium text-gray-100"
                  >
                    I Agree
                  </label>
                  <p id="comments-description" className="text-gray-400">
                    Creation of a new project will include downloading
                    dependencies from github.
                  </p>
                </div>
              </div>
            </div>
          </div>
          {/* <div className="px-8">
            <Stepper step={step} />
          </div> */}
          {/* {stdout.length > 0 && ( */}
          <div className="w-full max-w-[750px] overflow-x-auto flex flex-col overflow-y-scroll text-wrap h-[400px] max-h-[400px] text-xs bg-black text-white rounded-md px-4 py-8">
            {logs.map((line, idx) => (
              <pre key={idx}>{line}</pre>
            ))}
            <div ref={logsRef} />
          </div>
          {/* )} */}
          <div className="flex items-center justify-end mt-4  w-full mb-4">
            <button
              onClick={createProject}
              className="flex space-x-4 items-center  px-3 py-2 rounded-md text-gray-100 border border-gray-600 hover:bg-gray-700"
            >
              <div className="">
                {inProgress ? (
                  <div className="mr-2">
                    <Spinner />
                  </div>
                ) : (
                  <div className="mr-2">
                    <svg
                      className="w-3.5 h-3.5 text-white"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 16 12"
                    >
                      <path
                        stroke="currentColor"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="2"
                        d="M1 5.917 5.724 10.5 15 1.5"
                      />
                    </svg>
                  </div>
                )}
              </div>
              Create Project
            </button>
          </div>
        </div>
      </div>
    </>
  );
}
