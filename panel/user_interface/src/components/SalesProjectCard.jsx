export default function SalesProjectCard() {
  return (
    <div className="my-sales-project__content">
      <div className="p-3 rounded-lg flex flex-col gap-6 relative border border-gray-400">
        {/* My Sale Project Box Header */}
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <img src="images/githun.png" width="40" alt="" />
            <span className="text-base font-semibold">Github</span>
          </div>
          <div>
            <img
              width="20"
              src="images/delete.png"
              className="cursor-pointer select-none"
            />
          </div>
        </div>
        <div className="flex items-center justify-between">
          <span>Contact</span>
          <div className="bg-gray-400 flex items-center justify-between gap-2 rounded-full px-2">
            {/* <img
              src="images/man-avatar.jpg"
              className="align-middle rounded-full"
              width="16"
              alt=""
            /> */}
            <span>Amanda.s</span>
          </div>
        </div>
        <div className="flex justify-between items-center">
          <span>Amount</span>
          <span className="font-medium text-gray-700">$100K</span>
        </div>
      </div>
      <span className="divide"></span>
    </div>
  );
}
