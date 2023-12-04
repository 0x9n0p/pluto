'use client'

import { Content, Theme } from '@carbon/react';

import MainHeader from '@/components/MainHeader/MainHeader';

export function Providers({ children }) {
  return (
    <div>
        <Theme theme="g100">
          <MainHeader />
        </Theme>
        <Content>
          {children}
        </Content>
    </div>
  )
}
