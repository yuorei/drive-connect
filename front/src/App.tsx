import ButtomNavigation from './GlobalComponents/BottomNavigation.tsx'
import { Outlet } from 'react-router-dom'

function App() {

  return (
    <>
      <Outlet />
      <ButtomNavigation />
    </>
  )
}

export default App
