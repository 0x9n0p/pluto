export default function LogoutBtn() {
  return (
    <button className="bg-white rounded-lg text-sm cursor-pointer flex items-center justify-center gap-2 text-center text-gray-800 absolute w-4/5 py-[.6rem] px-4 outline-none border-none shadow-btn bottom-6 left-1/2 -translate-x-1/2">
      <svg className="w-5 h-5 text-blue-600">
        <use xlinkHref="#logout"></use>
      </svg>
      Log Out
    </button>
  );
}
