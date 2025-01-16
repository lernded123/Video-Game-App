import React from 'react'
import { createRootRoute, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import Navbar from '../components/Navbar/Navbar'

export const Route = createRootRoute({
  component: () => (
    <>
    <div>
        <Navbar />
      </div>
      <hr />
      <Outlet />
      <TanStackRouterDevtools />
    </>
  ),
})