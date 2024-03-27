import { EnvelopeIcon, PhotoIcon } from "@heroicons/react/20/solid";
import Heading from "../components/Heading";
import Container from "./Container";
import { useCallback, useState } from "react";
import { OpenFile } from "../../wailsjs/go/main/App";
// import { useApplictionStore } from "../store/store";
import { useDirDialog } from "../hooks/useFs";

export default function Feedback() {
  const [name, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [message, setMessage] = useState<string>("");

  const { filepath, getFilepath } = useDirDialog();
  // const sf = useApplictionStore((state) => state.submitFeedback);

  function submitFeedback() {
    const fb = {
      name,
      email,
      filepath,
      message,
    };

    // sf(fb);
  }

  return (
    <Container>
      <Heading
        title="Feedback"
        description="Share your feedback and ideas about using Tectone Desktop. We highly value and encourage feedback from the Tectone community."
      />
      <div className="w-full flex flex-col max-w-3xl space-y-6 ">
        <div className="flex justify-between space-x-4">
          <div className="flex-1">
            <label
              htmlFor="name"
              className="block text-sm font-medium leading-6 text-gray-300"
            >
              Name
            </label>
            <div className="relative mt-2 rounded-md shadow-sm">
              <input
                type="email"
                name="name"
                id="name"
                className="block w-full rounded-md border-0 py-1.5 pl-4 bg-gray-900 text-gray-300 focus:ring-gray-700 ring-1 ring-inset ring-gray-700 placeholder:text-gray-500 focus:ring-2 focus:ring-inset sm:text-sm sm:leading-6"
                placeholder="Firstname Lastname"
                value={name}
                onChange={(e) => setName(e.currentTarget.value)}
              />
            </div>
          </div>
          <div className="flex-1">
            <label
              htmlFor="email"
              className="block text-sm font-medium leading-6 text-gray-300"
            >
              Email
            </label>
            <div className="relative mt-2 rounded-md shadow-sm">
              <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                <EnvelopeIcon
                  className="h-5 w-5 text-gray-400"
                  aria-hidden="true"
                />
              </div>
              <input
                type="email"
                name="email"
                id="email"
                className="block w-full rounded-md border-0 py-1.5 pl-10 bg-gray-900 text-gray-300 ring-1 ring-inset ring-gray-700 placeholder:text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                placeholder="you@example.com"
                value={email}
                onChange={(e) => setEmail(e.currentTarget.value)}
              />
            </div>
          </div>
        </div>
        <div>
          <label
            htmlFor="file-upload"
            className="block text-sm font-medium leading-6 text-gray-300"
          >
            Screenshot(s)
          </label>
          <div className="mt-2 flex justify-center rounded-lg border border-dashed border-white/25 px-3 py-5">
            <div className="text-center">
              <PhotoIcon
                className="mx-auto h-12 w-12 text-gray-500"
                aria-hidden="true"
              />
              <div className="my-4 flex text-sm leading-6 text-gray-400">
                <button
                  id="file-upload"
                  onClick={getFilepath}
                  className="px-3 border border-gray-700 mx-auto rounded-md py-2 hover:bg-gray-700 hover:text-gray-300"
                >
                  Upload a File
                </button>
              </div>
              <p className="text-xs leading-5 text-gray-400">
                PNG, JPG, GIF up to 10MB
              </p>
            </div>
          </div>
        </div>
        <div>
          <label
            htmlFor="message"
            className="block text-sm font-medium leading-6 text-white"
          >
            Message
          </label>
          <div className="mt-2">
            <textarea
              id="message"
              name="about"
              rows={3}
              className="block w-full rounded-md border-0 bg-white/5 py-1.5 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-indigo-500 sm:text-sm sm:leading-6"
              defaultValue={""}
              value={message}
              onChange={(e) => setMessage}
            />
          </div>
          <p className="mt-1 text-xs leading-6 text-gray-500">
            Write a few sentences about any feedback you'd like to give.
          </p>
        </div>
        <div className="flex justify-end">
          <button
            onClick={submitFeedback}
            className="px-3 py-2 rounded-md border border-gray-700 text-gray-400 hover:bg-gray-800 hover:text-gray-300"
          >
            Submit Feedback
          </button>
        </div>
      </div>
    </Container>
  );
}
