// @flow
import classes from "./Navbar.module.scss";
import Link from "next/link";
const Navbar: React.FunctionComponent = () => {
  return (
    <nav
      className={`navbar navbar-expand-lg ${classes.root}`}
    >
      <div className={`container-fluid ${classes.navbarBody}`}>
        <Link href="/" as="/">
          <a className={`navbar-brand ${classes.navbarBrand}`} href="#">
            Desafioloca
          </a>
        </Link>
          <div
            className={`collapse navbar-collapse ${classes.navbarRightRoot}`}
            id="navbarSupportedContent"
          >
            <ul className={`navbar-nav ml-auto ${classes.navbarRightBody}`}>
              <li className={`nav-item ${classes.bankAccountInfo}`}>
                <img
                  src="/img/icon_user.png"
                  alt=""
                  className={classes.iconUser}
                />
              </li>
            </ul>
          </div>
      </div>
    </nav>
  );
};

export default Navbar;
