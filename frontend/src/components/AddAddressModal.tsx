import React, { useState, Fragment, useCallback } from "react";
import { Dialog, Transition } from "@headlessui/react";
import { model } from "../../wailsjs/go/models";
import { GenerateAddr } from "../../wailsjs/go/main/SDK";
interface AddAddressModalProps {
  open: boolean;
  setOpen: (open: boolean) => void;
  onSave: (alloc: model.Allocation) => void;
}

export default function AddAddressModal({
  open,
  setOpen,
  onSave,
}: AddAddressModalProps) {
  const [addr, setAddr] = useState<string>("");
  const [comment, setComment] = useState<string>("");
  const [algo, setAlgo] = useState<number>(0);
  const [onl, setOnl] = useState<number | undefined>();
  const [sel, setSel] = useState<string | undefined>();
  const [vote, setVote] = useState<string | undefined>();
  const [voteKD, setVoteKD] = useState<number | undefined>();
  const [voteLst, setVoteLst] = useState<number | undefined>();

  const getNewAddr = useCallback(async () => {
    const generated = await GenerateAddr();
    setAddr(generated);
  }, []);

  function save() {
    console.log({ algo });
    let alloc: model.Allocation = new model.Allocation({
      addr,
      comment,
      state: {
        algo,
        onl,
        sel,
        vote,
        voteKD,
        voteLst,
      },
    });

    console.log(alloc);
    onSave(alloc);
  }
  return (
    <>
      <Transition.Root show={open} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={setOpen}>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
          </Transition.Child>

          <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
            <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                enterTo="opacity-100 translate-y-0 sm:scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              >
                <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-gray-800 px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-xl sm:p-6">
                  <div className="flex flex-col space-y-2">
                    <div className="text-lg text-gray-400 text-center w-full mb-6">
                      Create new address
                    </div>
                    <div className="">
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
                        value={addr}
                        onChange={(e) => setAddr(e.target.value)}
                      />
                    </div>
                    <div className="">
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
                        value={comment}
                        onChange={(e) => setComment(e.target.value)}
                      />
                    </div>
                    <div className="flex w-full space-x-4">
                      <div className="flex-1">
                        <label
                          htmlFor="algo"
                          className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                        >
                          Algo
                        </label>
                        <input
                          type="number"
                          id="algo"
                          className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                          value={algo}
                          onChange={(e) =>
                            setAlgo((old) => parseInt(e.target.value))
                          }
                        />
                      </div>
                      <div className="flex-1">
                        <label
                          htmlFor="onl"
                          className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                        >
                          Onl
                        </label>
                        <input
                          type="number"
                          id="onl"
                          className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                          value={onl}
                          onChange={(e) =>
                            setOnl((old) => parseInt(e.target.value))
                          }
                        />
                      </div>
                    </div>
                    <div className="">
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
                        value={sel}
                        onChange={(e) => setSel(e.target.value)}
                      />
                    </div>
                    <div className="">
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
                        value={vote}
                        onChange={(e) => setVote(e.target.value)}
                      />
                    </div>
                    <div className="flex space-x-4 w-full">
                      <div className="flex-1">
                        <label
                          htmlFor="voteKD"
                          className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                        >
                          VoteKD
                        </label>
                        <input
                          type="number"
                          id="voteKD"
                          className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                          value={voteKD}
                          onChange={(e) =>
                            setVoteKD((old) => parseInt(e.target.value))
                          }
                        />
                      </div>
                      <div className="flex-1">
                        <label
                          htmlFor="voteLst"
                          className=" capitalize block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                        >
                          VoteLst
                        </label>
                        <input
                          type="number"
                          id="voteLst"
                          className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                          value={voteLst}
                          onChange={(e) =>
                            setVoteLst((old) => parseInt(e.target.value))
                          }
                        />
                      </div>
                    </div>
                  </div>
                  <div className="mt-5 sm:mt-6">
                    <button
                      type="button"
                      className="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                      onClick={() => {
                        save();
                        setOpen(false);
                      }}
                    >
                      Add new address
                    </button>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  );
}
