import {
  Header,
  HeaderContainer,
  HeaderName,
  HeaderNavigation,
  HeaderMenuButton,
  HeaderMenuItem,
  HeaderGlobalBar,
  HeaderGlobalAction,
  SkipToContent,
  SideNav,
  SideNavItems,
  HeaderSideNavItems,
  SideNavMenu,
  SideNavMenuItem,
  SideNavLink,
  StoryContent,
  HeaderMenu,
  SwitcherItem,
  HeaderPanel,
  SwitcherDivider,
  Switcher,
} from '@carbon/react';
import {
  Switcher as SwitcherIcon,
  UserAvatar as UserAvatarIcon,
} from '@carbon/icons-react';

import Link from 'next/link';

const MainHeader = () => (
  <HeaderContainer
    render={({ isSideNavExpanded, onClickSideNavExpand }) => (
      <Header aria-label="PlutoEngine">
        <SkipToContent />
        <HeaderMenuButton
          aria-label="Open menu"
          onClick={onClickSideNavExpand}
          isActive={isSideNavExpanded}
        />

        <Link href="/" passHref legacyBehavior>
          <HeaderName prefix="PlutoEngine">Panel</HeaderName>
        </Link>

        <HeaderNavigation aria-label="PlutoEngine">
          <Link href="/pipelines" passHref legacyBehavior>
            <HeaderMenuItem>Pipelines</HeaderMenuItem>
          </Link>
          <Link href="/logsview" passHref legacyBehavior>
            <HeaderMenuItem>Watch logs in real-time</HeaderMenuItem>
          </Link>
        </HeaderNavigation>

        <HeaderGlobalBar>
          {/*<HeaderGlobalAction*/}
          {/*  aria-label="Notifications"*/}
          {/*  tooltipAlignment="center"*/}
          {/*>*/}
          {/*  <Notification size={20} />*/}
          {/*</HeaderGlobalAction>*/}
          <HeaderGlobalAction aria-label="User" tooltipAlignment="center">
            <UserAvatarIcon size={20} />
          </HeaderGlobalAction>
          <HeaderGlobalAction
            aria-label={isSideNavExpanded ? 'Close switcher' : 'Open switcher'}
            aria-expanded={isSideNavExpanded}
            isActive={isSideNavExpanded}
            onClick={onClickSideNavExpand}
            tooltipAlignment="end"
            id="switcher-button"
          >
            <SwitcherIcon size={20} />
          </HeaderGlobalAction>
        </HeaderGlobalBar>

        <HeaderPanel
          expanded={isSideNavExpanded}
          onHeaderPanelFocus={onClickSideNavExpand}
          href="#switcher-button"
        >
          <Switcher aria-label="Switcher" expanded={isSideNavExpanded}>
            <SwitcherItem aria-label="Home" href="/">
              Home
            </SwitcherItem>
            <SwitcherItem aria-label="Pipelines" href="/pipelines">
              Pipelines
            </SwitcherItem>
            <SwitcherItem href="/logsview" aria-label="Watch logs in real-time">
              Watch logs in real-time
            </SwitcherItem>
            <SwitcherDivider />

            {/*TODO*/}
            {/*<SwitcherItem href='#' aria-label='Link 3'>*/}
            {/*  Settings*/}
            {/*</SwitcherItem>*/}
            {/*<SwitcherItem href='#' aria-label='Link 4'>*/}
            {/*  Backup*/}
            {/*</SwitcherItem>*/}
            {/*<SwitcherDivider />*/}

            <SwitcherItem
              href="https://plutoengine.ir/docs/"
              aria-label="Documentation"
            >
              Documentation
            </SwitcherItem>
            <SwitcherItem
              href="https://t.me/PlutoEngineCommunity"
              aria-label="Community"
            >
              Community
            </SwitcherItem>
            <SwitcherItem
              href="https://github.com/0x9n0p/pluto"
              aria-label="Source Code"
            >
              Source Code
            </SwitcherItem>
          </Switcher>
        </HeaderPanel>
      </Header>
    )}
  />
);
export const UnAuthenticatedHeader = () => (
  <HeaderContainer
    render={({ isSideNavExpanded, onClickSideNavExpand }) => (
      <Header aria-label="PlutoEngine">
        <SkipToContent />
        <HeaderMenuButton
          aria-label="Open menu"
          onClick={onClickSideNavExpand}
          isActive={isSideNavExpanded}
        />
        <Link href="/" passHref legacyBehavior>
          <HeaderName prefix="PlutoEngine">Panel</HeaderName>
        </Link>

        <HeaderGlobalBar>
          <HeaderGlobalAction
            aria-label="Notifications"
            tooltipAlignment="center"
          >
            <Notification size={20} />
          </HeaderGlobalAction>
          <HeaderGlobalAction
            aria-label="User Avatar"
            tooltipAlignment="center"
          >
            <UserAvatar size={20} />
          </HeaderGlobalAction>
          <HeaderGlobalAction aria-label="App Switcher" tooltipAlignment="end">
            <SwitcherIcon size={20} />
          </HeaderGlobalAction>
        </HeaderGlobalBar>
      </Header>
    )}
  />
);

export default MainHeader;
