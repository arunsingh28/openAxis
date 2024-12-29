import { RouteObject } from 'react-router-dom';

import { appPath } from '@/utils/path';

import IndexLayout from '@/layouts/index';

import Index from '@/views/index';

export const routes: RouteObject[] = [
  {
    path: appPath.DASHBOARD.INDEX,
    element: <IndexLayout />,
    children: [
      {
        index: true,
        element: <Index />
      }
    ]
  }
];
