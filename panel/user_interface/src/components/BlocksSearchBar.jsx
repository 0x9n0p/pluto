export default function BlocksSearchBar() {
  return (
    <div className="relative w-full">
      <input
        type="text"
        placeholder="Search Blocks"
        className="p-3 pl-10 w-full rounded bg-transparent text-base outline-none border border-gray-400"
      />
      <div className="absolute w-5 text-gray-600 left-3 top-1/2 -translate-y-1/2">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth="1.5"
          stroke="currentColor"
          className="w-6 h-6"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z"
          />
        </svg>
      </div>
    </div>
  );
}
