export default function AddPiplineBtn() {
  return (
    <button className="bg-white rounded-lg text-sm cursor-pointer flex items-center gap-1 py-[.6rem] px-4 outline-none border-none shadow-btn">
      <svg className="w-5 h-5 text-blue-600">
        <use xlinkHref="#plus"></use>
      </svg>
      New
    </button>
  );
}
