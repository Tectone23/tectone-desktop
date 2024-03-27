import { model } from "../../wailsjs/go/models";
import { Accordion } from "flowbite-react";
import AllocationFormEntry from "./AllocationFormEntry";
import { useEffect } from "react";

interface AllocationsFormProps {
  allocations: model.Allocation[];
  deleteAllocation: (index: number) => void;
  onChange: React.Dispatch<React.SetStateAction<model.Allocation[]>>;
}

export default function AllocationsForm({
  allocations,
  deleteAllocation,
  onChange,
}: AllocationsFormProps) {
  // console.log({ allocations });
  useEffect(() => {
    const evt = new Event("DOMContentLoaded", {
      bubbles: true,
      cancelable: false,
    });
    document.dispatchEvent(evt);
  }, [allocations]);

  function updateOnChange(index: number, al: model.Allocation) {
    const allocs = allocations.map((a, idx) => (idx == index ? al : a));
    onChange(allocs);
  }
  return (
    <>
      <Accordion collapseAll flush>
        {allocations &&
          allocations.map((allocation, index) => {
            return (
              <Accordion.Panel key={index}>
                <Accordion.Title className="py-6 px-4 text-xs">
                  {allocation.addr}
                </Accordion.Title>
                <Accordion.Content className="py-6 px-4 relative">
                  <AllocationFormEntry
                    onChange={updateOnChange}
                    onDelete={deleteAllocation}
                    index={index}
                    allocation={allocation}
                  />
                  {/* <div className="w-full right-4 top-4 absolute flex justify-end z-0">
                  <button
                    onClick={() => {
                      deleteAllocation(index);
                    }}
                  >
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
                        d="M6 18 18 6M6 6l12 12"
                      />
                    </svg>
                  </button>
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="addr"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Addr
                  </label>
                  <input
                    type="text"
                    id="addr"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.addr}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="comment"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Comment
                  </label>
                  <input
                    type="text"
                    id="comment"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.comment}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="algo"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Algo
                  </label>
                  <input
                    type="text"
                    id="algo"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.state.algo}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="onl"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Onl
                  </label>
                  <input
                    type="text"
                    id="onl"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.state.onl}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="sel"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Sel
                  </label>
                  <input
                    type="text"
                    id="sel"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.state.sel}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="vote"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    Vote
                  </label>
                  <input
                    type="text"
                    id="vote"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.state.vote}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="votekd"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    VoteKD
                  </label>
                  <input
                    type="text"
                    id="votekd"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.state.voteKD}
                  />
                </div>
                <div className="w-full h-full">
                  <label
                    htmlFor="votelst"
                    className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    VoteLst
                  </label>
                  <input
                    type="text"
                    id="votelst"
                    className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={allocation.state.voteLst}
                  />
                </div> */}
                  {/* {Object.entries(allocation).map(([key, value], index) => {
                  if (key == "state") {
                    return Object.entries(value).map(([k, v], i) => {
                      return (
                        <div className="w-full h-full" key={i}>
                          <label
                            htmlFor={k}
                            className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                          >
                            {k}
                          </label>
                          <input
                            type="text"
                            id={k}
                            className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                            value={v as any}
                          />
                        </div>
                      );
                    });
                  } else {
                    return (
                      <div className="" key={index}>
                        <label
                          htmlFor={key}
                          className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                        >
                          {key}
                        </label>
                        <input
                          type="text"
                          id={key}
                          className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                          value={value}
                        />
                      </div>
                    );
                  }
                })} */}
                  {/* <div className="flex w-full justify-end mt-6">
                  <button className="px-3 py-2 border border-gray-700 text-gray-300 rounded-lg hover:bg-gray-800">
                    Save
                  </button>
                </div> */}
                </Accordion.Content>
              </Accordion.Panel>
            );
          })}
      </Accordion>
    </>
  );
}
