import { Disclosure } from "@headlessui/react";
import { ChevronRightIcon } from "@heroicons/react/20/solid";
import AllocationItem from "./AllocationItem";
import { model } from "../../wailsjs/go/models";
import classNames from "classnames";
import { useEffect, useState } from "react";

interface AllocationsProps {
  allocations: model.Allocation[];
  deleteAllocation: (index: number) => void;
  onSave: (alloc: model.Allocation) => void;
}

export default function Allocations({
  allocations,
  deleteAllocation,
  onSave,
}: AllocationsProps) {
  const [alloc, setAlloc] = useState(allocations);

  useEffect(() => {
    setAlloc(allocations);
  }, [allocations]);
  return (
    <>
      {alloc &&
        alloc.map((allocation, index) => (
          <Disclosure>
            {({ open }) => (
              /* Use the `open` state to conditionally change the direction of an icon. */
              <div className="w-full flex flex-col relative mt-4">
                <Disclosure.Button className="w-full flex z-10 items-center  text-sm bg-gray-800 rounded-md border-gray-700 border px-2 py-2">
                  {allocation.addr}
                  <ChevronRightIcon
                    className={classNames(
                      open ? "rotate-90 transform " : "",
                      "text-white h-4 w-4 ml-3"
                    )}
                  />
                </Disclosure.Button>
                <Disclosure.Panel className="flex px-4 flex-col bg-gray-900 rounded-md mt-4 py-4">
                  <AllocationItem
                    allocation={allocation}
                    onSave={onSave}
                    onDelete={deleteAllocation}
                    index={index}
                  />
                </Disclosure.Panel>
              </div>
            )}
          </Disclosure>
        ))}
    </>
  );
}
