export default function SalesProjectCard() {
  return (
    <div className="my-sales-project__content">
      <div className="my-sales-project__box">
        {/* My Sale Project Box Header */}
        <div className="my-sales-project__box-header">
          <div className="my-sales-project__box-header--right">
            <img src="images/githun.png" width="40" alt="" />
            <span className="my-sales-project__box-title">Github</span>
          </div>
          <div>
            <img
              width="20"
              src="images/delete.png"
              className="my-sales-project__box-remove"
            />
          </div>
        </div>
        <div className="my-sales-project__box-contact">
          <span>Contact</span>
          <div className="my-sales-project__box-contact-person">
            <img
              src="images/man-avatar.jpg"
              className="my-sales-project__box-contact-image"
              width="16"
              alt=""
            />
            <span>Amanda.s</span>
          </div>
        </div>
        <div className="my-sales-project__box-amount">
          <span>Amount</span>
          <span className="my-sales-project__box-number">$100K</span>
        </div>
      </div>
      <span className="divide"></span>
    </div>
  );
}
