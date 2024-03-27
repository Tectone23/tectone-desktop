interface StepperProps {
  step: number;
}
export default function Stepper({ step }: StepperProps) {
  return (
    <ol className="relative text-gray-500 border-s border-gray-200 dark:border-gray-700 dark:text-gray-400">
      <li className="mb-10 ms-6">
        <span className="absolute flex items-center justify-center w-8 h-8 border border-white bg-gray-800 rounded-full -start-4 ring-4 ring-white dark:ring-gray-900 ">
          <svg
            className={`"w-3.5 h-3.5 " ${step >= 1 ? "text-white" : "hidden"}`}
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
        </span>
        <h3 className="font-medium leading-tight">Algorand Go SDK</h3>
        <p className="text-sm">Fetching git repository from Github</p>
      </li>
      <li className="mb-10 ms-6">
        <span className="absolute flex items-center justify-center w-8 h-8 border border-white rounded-full -start-4   bg-gray-800">
          <svg
            className={`"w-3.5 h-3.5 " ${step >= 1 ? "text-white" : "hidden"}`}
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
        </span>
        <h3 className="font-medium leading-tight">Algorand Sandbox</h3>
        <p className="text-sm">fetching and setting up sandbox</p>
      </li>
      {/* <li className="mb-10 ms-6">
        <span className="absolute flex items-center justify-center w-8 h-8 border border-white bg-gray-800 rounded-full -start-4  ">
          <svg
            className={`"w-3.5 h-3.5 " ${step >= 3 ? "text-white" : "hidden"}`}
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            viewBox="0 0 18 20"
          >
            <path d="M16 1h-3.278A1.992 1.992 0 0 0 11 0H7a1.993 1.993 0 0 0-1.722 1H2a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2Zm-3 14H5a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2Zm0-4H5a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2Zm0-5H5a1 1 0 0 1 0-2h2V2h4v2h2a1 1 0 1 1 0 2Z" />
          </svg>
        </span>
        <h3 className="font-medium leading-tight">Review</h3>
        <p className="text-sm">Step details here</p>
      </li> */}
    </ol>
  );
}
