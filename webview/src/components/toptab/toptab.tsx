import { Button, Navbar, NavbarContent, NavbarItem } from "@heroui/react";
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { IoSettingsSharp, IoTerminal } from "react-icons/io5";
import "./toptab.css";
import { ThemeSwitcher } from '../themeswitcher/themeswitcher';
const pageList = [
  { path: 'devices', label: 'Devices' },
  { path: 'telemetryservices', label: 'Telemetry Services' },
];

export default function TopTab() {
  const location = useLocation();
  const pathname = location.pathname.split('/')[1];
  const navigate = useNavigate();

  return (
    <Navbar classNames={{ wrapper: "px-0 h-auto" }} >
      <NavbarContent className="sm:flex" justify="center">
        <div className="button-group">
          {pageList.map((item, index) => (
            <Button
              key={index}
              as={Link}
              to={`/${item.path}`}
              color={pathname === item.path ? "primary" : "default"}
              radius="none"
              className={`group-button ${index === 0 ? 'first-button' : ''} ${index === pageList.length - 1 ? 'last-button' : ''
                }`}
              variant={pathname === item.path ? "solid" : "flat"}
            >
              <p className='text-lg'>{item.label}</p>
            </Button>
          ))}
        </div>
      </NavbarContent>
      <NavbarContent className="sm:flex" justify="end">
        <ThemeSwitcher />
        <Button variant='light' isIconOnly disableRipple disableAnimation onPress={() => navigate("/terminal")}>
          <IoTerminal className='h-full w-auto p-2' />
        </Button>
        <Button variant='light' isIconOnly disableRipple disableAnimation className='button-setting' onPress={() => navigate("/settings")}>
          <IoSettingsSharp className='h-full w-auto' />
        </Button>
      </NavbarContent>
    </Navbar>
  );
}

