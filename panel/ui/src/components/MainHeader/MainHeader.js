'use client';

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
} from '@carbon/react';
import {
  Switcher,
  Notification,
  UserAvatar,
} from '@carbon/icons-react';

import Link from 'next/link';

const MainHeader = () => (
  <HeaderContainer
    render={({ isSideNavExpanded, onClickSideNavExpand }) => (
      <Header aria-label='PlutoEngine'>
        <SkipToContent />
        <HeaderMenuButton
          aria-label='Open menu'
          onClick={onClickSideNavExpand}
          isActive={isSideNavExpanded}
        />
        <Link href='/' passHref legacyBehavior>
          <HeaderName prefix='PlutoEngine'>
            Panel
          </HeaderName>
        </Link>
        <HeaderNavigation aria-label='PlutoEngine'>
          <Link href='/pipelines/create' passHref legacyBehavior>
            <HeaderMenuItem>
              Create a new pipeline
            </HeaderMenuItem>
          </Link>
        </HeaderNavigation>
        <SideNav
          aria-label='Side navigation'
          expanded={isSideNavExpanded}
          isPersistent={false}>
          <SideNavItems>
            <HeaderSideNavItems>
              <Link href='/pipelines/create' passHref legacyBehavior>
                <HeaderMenuItem>
                  Create a new pipeline
                </HeaderMenuItem>
              </Link>
            </HeaderSideNavItems>
          </SideNavItems>
        </SideNav>
        <HeaderGlobalBar>
          <HeaderGlobalAction
            aria-label='Notifications'
            tooltipAlignment='center'>
            <Notification size={20} />
          </HeaderGlobalAction>
          <HeaderGlobalAction
            aria-label='User Avatar'
            tooltipAlignment='center'>
            <UserAvatar size={20} />
          </HeaderGlobalAction>
          <HeaderGlobalAction aria-label='App Switcher' tooltipAlignment='end'>
            <Switcher size={20} />
          </HeaderGlobalAction>
        </HeaderGlobalBar>
      </Header>
    )}
  />
);

export default MainHeader;
