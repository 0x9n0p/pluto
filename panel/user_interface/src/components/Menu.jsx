export default function Menu() {
  return (
    <ul className="flex flex-col gap-12 py-0 px-2">
      <li className="sidebar__menu-item">
        <a href="#" className="sidebar__menu-link">
          <svg className="sidebar__menu-icon">
            <use xlinkHref="#phone"></use>
          </svg>
          Contact
        </a>
      </li>
      <li className="sidebar__menu-item sidebar__menu-item--active">
        <a className="sidebar__menu-link" href="#">
          <svg className="sidebar__menu-icon">
            <use xlinkHref="#document-text"></use>
          </svg>
          Piplines
        </a>
      </li>
      <li className="sidebar__menu-item">
        <a className="sidebar__menu-link" href="#">
          <svg className="sidebar__menu-icon">
            <use xlinkHref="#user"></use>
          </svg>
          Admin
        </a>
      </li>
      <li className="sidebar__menu-item">
        <a className="sidebar__menu-link" href="#">
          <svg className="sidebar__menu-icon">
            <use xlinkHref="#setting"></use>
          </svg>
          Setting
        </a>
      </li>
    </ul>
  );
}
