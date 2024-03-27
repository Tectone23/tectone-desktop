import { useState } from "react";

export default function DispenserForm() {
  const [addr, setAddr] = useState<string>("");
  const [amount, setAmount] = useState<number>(0);

  function saveForm() {
    return;
  }
  return (
    <>
      <div className=" flex flex-col space-y-4 w-full max-w-2xl max-h-full">
        <div>
          <label
            htmlFor="addr"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Address
          </label>
          <input
            type="text"
            id="addr"
            className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            value={addr}
          />
        </div>
        <div>
          <label
            htmlFor="amount"
            className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            Amount
          </label>
          <input
            type="text"
            id="amount"
            className="block w-full p-2 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-xs focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            value={amount}
          />
        </div>
        <div className="mt-8 w-full flex justify-end">
          <button
            className="px-3 py-2 border border-gray-700 rounded-lg hover:bg-gray-700"
            onClick={() => saveForm()}
          >
            Save
          </button>
        </div>
      </div>
    </>
  );
}
