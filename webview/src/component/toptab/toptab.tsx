import { Button, Navbar, NavbarContent, NavbarItem } from '@nextui-org/react';
import { Link, useLocation } from 'react-router-dom';
import "./toptab.css";

const pageList = [
  { path: 'devices', label: 'Devices' },
  { path: 'telemetryservices', label: 'Telemetry Services' },
];

export default function TopTab() {
  const location = useLocation();
  const pathname = location.pathname.split('/')[1];

  return (
    <div>
      <Navbar>
        <NavbarContent className="hidden sm:flex" justify="center">
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
      </Navbar>
    </div>
  );
}
