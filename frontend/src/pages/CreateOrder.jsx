import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { createOrder } from "../services/orderService";

function CreateOrder() {

  const navigate = useNavigate();

  const [provinces, setProvinces] = useState([]);
  const [originCities, setOriginCities] = useState([]);
  const [destinationCities, setDestinationCities] = useState([]);

  const [form, setForm] = useState({
    user_id: 1,

    sender_name: "",
    sender_phone: "",
    sender_address: "",

    receiver_name: "",
    receiver_phone: "",
    receiver_address: "",

    item_name: "",
    item_type: "",

    weight_kg: "",

    origin_city: "",
    destination_city: "",

    service_type: "EZ",

    payment_method: "TRANSFER"
  });

  useEffect(() => {
    fetch(
      "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
    )
      .then((res) => res.json())
      .then((data) => setProvinces(data))
      .catch((err) => console.error(err));
  }, []);

  const getOriginCities = async (provinceId) => {

    const response = await fetch(
      `https://www.emsifa.com/api-wilayah-indonesia/api/regencies/${provinceId}.json`
    );

    const data = await response.json();

    setOriginCities(data);
  };

  const getDestinationCities = async (
    provinceId
  ) => {

    const response = await fetch(
      `https://www.emsifa.com/api-wilayah-indonesia/api/regencies/${provinceId}.json`
    );

    const data = await response.json();

    setDestinationCities(data);
  };

  const handleChange = (e) => {

    setForm({
      ...form,
      [e.target.name]: e.target.value
    });

  };

  const handleSubmit = async (e) => {

    e.preventDefault();

    try {

      console.log(form);

      await createOrder(form);

      alert("Order berhasil dibuat");

      navigate("/orders");

    } catch (error) {

      console.error(error);

      alert("Gagal membuat order");

    }

  };

  return (

    <div className="form-container">

      <h1>Create Order</h1>

      <form onSubmit={handleSubmit}>

        <h3>Sender Information</h3>

        <input
          type="text"
          name="sender_name"
          placeholder="Sender Name"
          onChange={handleChange}
          required
        />

        <input
          type="text"
          name="sender_phone"
          placeholder="Sender Phone"
          onChange={handleChange}
          required
        />

        <textarea
          name="sender_address"
          placeholder="Sender Address"
          onChange={handleChange}
          required
        />

        <h3>Receiver Information</h3>

        <input
          type="text"
          name="receiver_name"
          placeholder="Receiver Name"
          onChange={handleChange}
          required
        />

        <input
          type="text"
          name="receiver_phone"
          placeholder="Receiver Phone"
          onChange={handleChange}
          required
        />

        <textarea
          name="receiver_address"
          placeholder="Receiver Address"
          onChange={handleChange}
          required
        />

        <h3>Package Information</h3>

        <input
          type="text"
          name="item_name"
          placeholder="Item Name"
          onChange={handleChange}
          required
        />

        <input
          type="text"
          name="item_type"
          placeholder="Item Type"
          onChange={handleChange}
          required
        />

        <input
          type="number"
          step="0.1"
          name="weight_kg"
          placeholder="Weight (Kg)"
          onChange={handleChange}
          required
        />

        <h3>Shipping Information</h3>

        <label>Origin Province</label>

        <select
          onChange={(e) =>
            getOriginCities(e.target.value)
          }
          required
        >
          <option value="">
            Select Province
          </option>

          {provinces.map((province) => (
            <option
              key={province.id}
              value={province.id}
            >
              {province.name}
            </option>
          ))}
        </select>

        <label>Origin City</label>

        <select
          name="origin_city"
          onChange={handleChange}
          required
        >
          <option value="">
            Select City
          </option>

          {originCities.map((city) => (
            <option
              key={city.id}
              value={city.name}
            >
              {city.name}
            </option>
          ))}
        </select>

        <label>Destination Province</label>

        <select
          onChange={(e) =>
            getDestinationCities(
              e.target.value
            )
          }
          required
        >
          <option value="">
            Select Province
          </option>

          {provinces.map((province) => (
            <option
              key={province.id}
              value={province.id}
            >
              {province.name}
            </option>
          ))}
        </select>

        <label>Destination City</label>

        <select
          name="destination_city"
          onChange={handleChange}
          required
        >
          <option value="">
            Select City
          </option>

          {destinationCities.map((city) => (
            <option
              key={city.id}
              value={city.name}
            >
              {city.name}
            </option>
          ))}
        </select>

        <label>Service Type</label>

        <select
          name="service_type"
          onChange={handleChange}
        >
          <option value="EZ">
            EZ Regular
          </option>

          <option value="DOC">
            DOC Document
          </option>

          <option value="JSD">
            JSD Same Day
          </option>

          <option value="JND">
            JND Next Day
          </option>
        </select>
        <h3>Payment Information</h3>

        <label>Payment Method</label>
        <select
          name="payment_method"
          value={form.payment_method}
          onChange={handleChange}
          required
        >
          <option value="TRANSFER">
            VA BCA
          </option>

          <option value="QRIS">
            QRIS
          </option>

          <option value="EWALLET">
            VA BRI
          </option>

          <option value="COD">
            Cash On Delivery
          </option>
        </select>
        <button
          type="submit"
          className="btn-primary"
        >
          Create Order
        </button>

      </form>

    </div>

  );
}

export default CreateOrder;