import { Routes, Route } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import Orders from "./pages/Orders";
import CreateOrder from "./pages/CreateOrder";
import Tracking from "./pages/Tracking";
import Profile from "./pages/Profile";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/orders" element={<Orders />} />
      <Route path="/orders/create" element={<CreateOrder />} />
      <Route path="/tracking" element={<Tracking />} />
      <Route path="/profile" element={<Profile />} />
    </Routes>
  );
}

export default App;