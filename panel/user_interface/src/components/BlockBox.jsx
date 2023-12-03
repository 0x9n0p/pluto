export default function BlocksBox({ icon, title, description }) {
  return (
    <div
      className="flex gap-4 items-center"
      draggable="true"
      id="blocks__box--1"
    >
      <div className="bg-gray-400 w-10 h-10 flex items-center justify-center rounded-lg px-0 py-4">
        <span className="text-gray-50 text-xs">
          <img src={icon} alt="Icon" />
        </span>
      </div>
      <div className="flex flex-col gap-2">
        <h1 className="text-xl font-semibold">{title}</h1>
        <p className="text-gray-500">{description}</p>
      </div>
    </div>
  );
}
