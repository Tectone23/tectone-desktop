import { Fragment, useEffect, useState } from "react";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon, ChevronUpDownIcon } from "@heroicons/react/20/solid";
import { main } from "../../wailsjs/go/models";

function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

interface ProjectSelectProps {
  projects: main.Project[];
}

export default function ProjectSelect({ projects }: ProjectSelectProps) {
  const [selected, setSelected] = useState<main.Project | null>(null);

  useEffect(() => {
    if (!projects) return;
    setSelected(projects[0]);
  }, [projects]);

  return (
    projects && (
      <Listbox value={selected} onChange={setSelected}>
        {({ open }) => (
          <>
            <Listbox.Label className="block text-sm font-medium leading-6 text-gray-200">
              Select a Project
            </Listbox.Label>
            <div className="relative mt-2">
              <Listbox.Button className="relative w-full cursor-default rounded-md  py-1.5 pl-3 pr-10 text-left text-gray-400 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6">
                <span className="block truncate">{selected?.name}</span>
                <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                  <ChevronUpDownIcon
                    className="h-5 w-5 text-gray-400"
                    aria-hidden="true"
                  />
                </span>
              </Listbox.Button>

              <Transition
                show={open}
                as={Fragment}
                leave="transition ease-in duration-100"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
              >
                <Listbox.Options className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md  text-gray-400 py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                  <Listbox.Option
                    key="default"
                    className=""
                    value="Select a project"
                  >
                    <span>Select a PRoject</span>
                  </Listbox.Option>
                  {projects.map((project, index) => (
                    <Listbox.Option
                      key={index}
                      className={({ active }) =>
                        classNames(
                          active
                            ? "bg-gray-200 text-gray-800"
                            : "text-gray-400",
                          "relative cursor-default select-none py-2 pl-3 pr-9"
                        )
                      }
                      value={project.name}
                    >
                      {({ selected, active }) => (
                        <>
                          <span
                            className={classNames(
                              selected ? "font-semibold" : "font-normal",
                              "block truncate"
                            )}
                          >
                            {project.name}
                          </span>

                          {selected ? (
                            <span
                              className={classNames(
                                active ? "text-white" : "text-indigo-600",
                                "absolute inset-y-0 right-0 flex items-center pr-4"
                              )}
                            >
                              <CheckIcon
                                className="h-5 w-5"
                                aria-hidden="true"
                              />
                            </span>
                          ) : null}
                        </>
                      )}
                    </Listbox.Option>
                  ))}
                </Listbox.Options>
              </Transition>
            </div>
          </>
        )}
      </Listbox>
    )
  );
}
