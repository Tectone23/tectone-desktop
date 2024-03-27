interface HeadingProps {
  title: string;
  description: string;
}
export default function Heading({ title, description }: HeadingProps) {
  return (
    <div className="flex flex-col max-w-2xl space-y-2">
      <div className="text-md capitalize font-semibold">{title}</div>
      <div className="text-sm text-gray-300">{description}</div>
    </div>
  );
}
