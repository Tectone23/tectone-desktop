import { useEffect, useState } from "react";
import { model } from "../../wailsjs/go/models";

interface AllocationFormEntryProps {
  allocation: model.Allocation;
  onChange: (index: number, alloc: model.Allocation) => void;
  onDelete: (idx: number) => void;
  index: number;
}
export default function AllocationFormEntry({
  allocation,
  onChange,
  onDelete,
  index,
}: AllocationFormEntryProps) {
  const [alloc, setAlloc] = useState(allocation);
  const [addr, setAddr] = useState<string>(alloc.addr);
  const [comment, setComment] = useState<string>(alloc.comment);
  const [algo, setAlgo] = useState<number>(alloc.state.algo);
  const [onl, setOnl] = useState<number | undefined>(alloc.state.onl);
  const [sel, setSel] = useState<string | undefined>(alloc.state.sel);
  const [vote, setVote] = useState<string | undefined>(alloc.state.vote);
  const [voteKD, setVoteKD] = useState<number | undefined>(
    allocation.state.voteKD
  );
  const [voteLst, setVoteLst] = useState<number | undefined>(
    allocation.state.voteLst
  );

  useEffect(() => {
    setAddr(allocation.addr);
    setComment(allocation.comment);
    setAlgo(allocation.state.algo);
    setAlgo(allocation.state.algo);
    setOnl(allocation.state.onl);
    setSel(allocation.state.sel);
    setVote(allocation.state.vote);
    setVoteKD(allocation.state.voteKD);
    setVoteLst(allocation.state.voteLst);
  }, [allocation]);

  useEffect(() => {
    var a: model.Allocation;
    a = {
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
      convertValues: alloc.convertValues,
    };
    onChange(index, a);
  }, [addr, comment, algo, onl, sel, vote, voteKD, voteLst]);

  return (
    <div className="flex flex-col space-y-2">
      <div className="w-full right-4 top-4 absolute flex justify-end z-0 text-gray-500 hover:text-gray-200">
        <button
          onClick={() => {
            onDelete(index);
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
          defaultValue={addr}
          // onChange={(e) => setAddr(e.currentTarget.value)}
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
          value={comment}
          onChange={(e) => setComment(e.currentTarget.value)}
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
          value={algo}
          onChange={(e) => setAlgo(e.currentTarget.value as unknown as number)}
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
          value={onl}
          onChange={(e) => setOnl(e.currentTarget.value as unknown as number)}
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
          value={sel}
          onChange={(e) => setSel(e.currentTarget.value)}
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
          value={vote}
          onChange={(e) => setVote(e.currentTarget.value)}
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
          value={voteKD}
          onChange={(e) =>
            setVoteKD(e.currentTarget.value as unknown as number)
          }
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
          value={voteLst}
          onChange={(e) =>
            setVoteLst(e.currentTarget.value as unknown as number)
          }
        />
      </div>
    </div>
  );
}
