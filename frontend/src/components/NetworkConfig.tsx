import { useEffect, useState } from "react";
import { main, model } from "../../wailsjs/go/models";
import AddAddressModal from "./AddAddressModal";
import AllocationsForm from "./AllocationsForm";
// import { useApplictionStore } from "../store/store";
import { useProject } from "../hooks/useProjects";
import Allocations from "./Allocation";
import AllocationDeletionConfirmationModal from "./AllocationDeletionConfirmationModal";

interface NetworkConfigProps {
  genesis: model.Genesis;
  onSave: (genesis: model.Genesis) => void;
}
export default function NetworkConfig({ genesis, onSave }: NetworkConfigProps) {
  const [genesisFields, setGenesisFields] = useState<model.Genesis>(genesis);
  const [id, setId] = useState(genesis?.id);
  const [fees, setFees] = useState(genesis?.fees);
  const [network, setNetwork] = useState(genesis?.network);
  const [proto, setProto] = useState(genesis?.proto);
  const [rwd, setRwd] = useState(genesis?.rwd);
  const [timestamp, setTimestamp] = useState(genesis?.timestamp);
  const [alloc, setAlloc] = useState(genesis?.alloc);
  const [open, setOpen] = useState<boolean>(false);
  const [filterBy, setFilterBy] = useState<keyof model.Allocation>("addr");

  // const project = useApplictionStore((state) => state.project);
  // const { project } = useProject()
  // const onSave = useApplictionStore((state) => state.saveTestNet);
  const [deletionIndex, setDeletionIndex] = useState<number | null>(null);
  const [showDeletionConfirmationModal, setShowDeletionConfirmationModal] =
    useState<boolean>(false);

  function deleteAllocation(index: number) {
    if (!genesisFields) return;
    var gf: model.Genesis;
    gf = {
      ...genesisFields,
      convertValues: genesis.convertValues,
      alloc: [...genesisFields.alloc],
    };

    const deleted = gf.alloc.splice(index, 1);
    onSave(gf);
    // onSave(project!!, gf);
    setGenesisFields(gf);
    setAlloc(gf.alloc);
    setOpen(false);
    setDeletionIndex(null);
  }

  function confirmDeletion(index: number) {
    setDeletionIndex(index);
    setShowDeletionConfirmationModal(true);
  }

  function saveGenesis() {
    if (!genesisFields) return;
    var g: model.Genesis = new model.Genesis({
      id,
      fees,
      network,
      proto,
      rwd,
      timestamp,
      alloc,
    });

    onSave(g);
    setGenesisFields(g);
  }

  function addAllocation(a: model.Allocation) {
    if (!genesisFields) return;
    var al: model.Allocation[] = genesisFields.alloc.map(
      (a) => new model.Allocation(a)
    );
    al.push(new model.Allocation(a));
    var gf: model.Genesis = new model.Genesis(genesisFields);

    gf.alloc = al;
    onSave(gf);
    // onSave(project!!, gf);
    setGenesisFields(gf);
    setAlloc(al);
  }

  useEffect(() => {
    setGenesisFields((old) => {
      const n = { ...genesis, convertValues: genesis.convertValues };
      return n;
    });
  }, [genesis]);

  return (
    <>
      <div className="relative w-full h-full">
        <button
          className="fixed z-50 bottom-[40px] right-10  px-4 py-2 rounded-md border border-gray-700 text-gray-300 hover:bg-gray-700 hover:border-gray-600"
          onClick={() => {
            saveGenesis();
          }}
        >
          Save
        </button>
      </div>
      <div className="flex space-y-4 w-full max-h-full z-0">
        <div className="flex flex-col w-full max-w-3xl ">
          <div>
            <label
              htmlFor="id"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              ID
            </label>
            <input
              type="text"
              id="id"
              className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={id}
              onChange={(e) => setId(e.currentTarget.value)}
            />
          </div>
          <div>
            <label
              htmlFor="network"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Network
            </label>
            <input
              type="text"
              id="id"
              className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={network}
              onChange={(e) => setNetwork(e.currentTarget.value)}
            />
          </div>
          <div>
            <label
              htmlFor="fees"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Fees
            </label>
            <input
              type="text"
              id="fees"
              className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={fees}
              onChange={(e) => setFees(e.currentTarget.value)}
            />
          </div>
          <div>
            <label
              htmlFor="proto"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Proto
            </label>
            <input
              type="text"
              id="proto"
              className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={proto}
              onChange={(e) => setProto(e.currentTarget.value)}
            />
          </div>
          <div>
            <label
              htmlFor="rwd"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              RWD
            </label>
            <input
              type="text"
              id="rwd"
              className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={rwd}
              onChange={(e) => setRwd(e.currentTarget.value)}
            />
          </div>
          <div>
            <label
              htmlFor="timestamp"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Timestamp
            </label>
            <input
              type="number"
              id="timestamp"
              className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              value={timestamp}
              onChange={(e) =>
                setTimestamp(e.currentTarget.value as unknown as number)
              }
              disabled
            />
          </div>

          <div className="flex flex-col mt-6 space-y-8">
            <div className="flex justify-between mt-8 mb-4 ">
              <h2 className="text-lg font-bold">Allocations: </h2>
              <button
                onClick={() => setOpen(true)}
                className="px-3 py-1 rounded-md bg-gray-800 border-gray-500 border"
              >
                ADD
              </button>
            </div>
            <div className="flex space-x-4">
              <div className="flex-1">
                <input
                  type="text"
                  className="w-full px-3 py-2 bg-gray-900 rounded-md border border-gray-700 text-gray-300 placeholder:text-gray-500 placeholder:italic"
                  placeholder="filter allocations..."
                  onChange={(e) => {
                    const result = genesis.alloc.filter((a: model.Allocation) =>
                      a[filterBy].toString().includes(e.currentTarget.value)
                    );
                    setAlloc(result);
                  }}
                />
              </div>
              <div className="py-1">
                <select
                  name="filterBy"
                  id="filterBy"
                  className="bg-gray-900 px-4 py-2 h-full text-sm"
                  onChange={(e) => {
                    setFilterBy(
                      e.currentTarget.value as keyof model.Allocation
                    );
                  }}
                  value={filterBy}
                >
                  <option value="addr">Addr</option>
                  <option value="comment" selected>
                    Comment
                  </option>
                  <option value="onl">Onl</option>
                  <option value="sel">Sel</option>
                  <option value="vote">Vote</option>
                  <option value="voteKD">VoteKD</option>
                  <option value="voteLst">VoteLst</option>
                </select>
              </div>
            </div>
            <AddAddressModal
              open={open}
              setOpen={setOpen}
              onSave={addAllocation}
            />

            <AllocationDeletionConfirmationModal
              open={showDeletionConfirmationModal}
              setOpen={setShowDeletionConfirmationModal}
              onDelete={deleteAllocation}
              index={deletionIndex}
            />

            {alloc && (
              <AllocationsForm
                allocations={alloc}
                deleteAllocation={confirmDeletion}
                onChange={setAlloc}
              />
              // <Allocations
              //   allocations={alloc}
              //   deleteAllocation={deleteAllocation}
              //   onSave={saveAllocation}
              // />
            )}
          </div>
        </div>
      </div>
    </>
  );
}
