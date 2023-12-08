`use client`;

import { Content, Theme } from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';

export default function LandingPage() {
  return (
    <>
      {localStorage.getItem('token') ? (
        <div>
          <Theme theme="g100">
            <MainHeader />
          </Theme>
          <Content>
            <div>LANDING PAGE</div>
          </Content>
        </div>
      ) : (
        window.location.assign('/auth')
      )}
    </>
  );
}
