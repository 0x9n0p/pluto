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
  SwitcherItem,
  HeaderPanel,
  SwitcherDivider,
  Switcher,
  OverflowMenuItem,
  OverflowMenu,
} from '@carbon/react';
import {
  Switcher as SwitcherIcon,
  UserAvatar as UserAvatarIcon,
} from '@carbon/icons-react';

import Link from 'next/link';
import axios from 'axios';
import { Address } from '@/settings';
import React from 'react';

const MainHeader = () => (
  <HeaderContainer
    render={({ isSideNavExpanded, onClickSideNavExpand }) => (
      <Header aria-label="PlutoEngine">
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
          <HeaderGlobalAction tooltipAlignment="hide">
            <OverflowMenu
              size={'lg'}
              renderIcon={UserAvatarIcon}
              // flipped={document?.dir === 'rtl'}
              menuOffset={{ left: -60 }}
            >
              <OverflowMenuItem itemText="Add a new admin" disabled />
              <OverflowMenuItem itemText="Change password" disabled />
              <OverflowMenuItem
                hasDivider
                isDelete
                itemText="Logout"
                onClick={(e) => {
                  axios
                    .post(Address() + '/api/v1/logout', {
                      headers: {
                        Authorization: `Bearer ${localStorage.getItem(
                          'token'
                        )}`,
                      },
                    })
                    .then(function (response) {
                      localStorage.removeItem('email');
                      localStorage.removeItem('token');
                      if (response.status === 200) {
                        window.location.href = '/auth';
                      }
                    })
                    .catch(function (error) {
                      localStorage.removeItem('email');
                      localStorage.removeItem('token');
                      if (error.response) {
                        if (error.response.status === 401) {
                          window.location.href = '/auth';
                          return;
                        }
                        window.location.href = '/auth';
                        console.error(error);
                      }
                    });
                }}
              />
            </OverflowMenu>
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
              href="https://github.com/0x9n0p/pluto/wiki"
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
        <Link href="/" passHref legacyBehavior>
          <HeaderName prefix="PlutoEngine">Panel</HeaderName>
        </Link>

        <HeaderGlobalBar>
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
            <SwitcherItem
              href="https://github.com/0x9n0p/pluto/wiki"
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

export default MainHeader;
