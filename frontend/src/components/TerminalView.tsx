interface TerminalViewProps {
  // children: React.ReactNode;
  logs: string[];
}
export default function TerminalView({ logs }: TerminalViewProps) {
  // get the logs from the terminal

  //
  return (
    <>
      <div className="flex flex-col max-h-full w-full rounded-md bg-black p-4 text-gray-200">
        {logs.map((log, index) => {
          return <pre key={index}>{log}</pre>;
        })}
      </div>
    </>
  );
}
