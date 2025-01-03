import { Button, Navbar, NavbarContent, NavbarItem } from '@nextui-org/react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { IoSettingsSharp } from "react-icons/io5";
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
    <div>
      <Navbar classNames={{ wrapper: "px-0 py-0" }}>
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
                {item.label}
              </Button>
            ))}
          </div>
        </NavbarContent>
        <NavbarContent className="sm:flex" justify="end">
          <ThemeSwitcher />
          <Button variant='light' isIconOnly disableRipple disableAnimation className='button-setting' onClick={() => navigate("/settings")}>
            <IoSettingsSharp className='h-full w-full' />
          </Button>
        </NavbarContent>
      </Navbar>
    </div>
  );
}

