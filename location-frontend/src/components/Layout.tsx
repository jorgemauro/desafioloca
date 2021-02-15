// @flow
import * as React from "react";
import { Account } from "../model";
import Footer from "./Footer";
import MainContent from "./MainContent";
import Navbar from "./Navbar";
interface LayoutProps {
  Account?: Account;
}
const Layout: React.FunctionComponent<LayoutProps> = (props) => {
  return (
    <>
      <Navbar />
      <MainContent>{props.children}</MainContent>
      <Footer />
    </>
  );
};

export default Layout;
