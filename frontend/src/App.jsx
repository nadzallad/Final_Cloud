import { Routes, Route } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import Orders from "./pages/Orders";
import CreateOrder from "./pages/CreateOrder";
import Tracking from "./pages/Tracking";
import Profile from "./pages/Profile";
import Payment from "./pages/Payment";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/orders" element={<Orders />} />
      <Route path="/orders/create" element={<CreateOrder />} />
      <Route path="/tracking" element={<Tracking />} />
      <Route path="/profile" element={<Profile />} />
      <Route path="/payment" element={<Payment />}/>
    </Routes>
  );
}

export default App;