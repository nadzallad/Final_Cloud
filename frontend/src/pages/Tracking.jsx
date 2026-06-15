import { useState } from "react";
import "./Tracking.css";
import trackingService from "../services/trackingService";

function Tracking() {
  const [formData, setFormData] = useState({
    tracking_id: "",
    no_resi: "",
    status: "",
    location: "",
    note: "",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {

      await trackingService.createTracking(
        formData
      );

      alert("Tracking berhasil disimpan!");

      setFormData({
        tracking_id: "",
        no_resi: "",
        status: "",
        location: "",
        note: "",
      });

    } catch (error) {
      alert("Gagal menyimpan data");
      console.error(error);
    }
  };

  return (
    <div className="tracking-container">
      <div className="tracking-card">
        <h1>🚚 Tracking Paket</h1>
        <p>Input informasi tracking pengiriman paket</p>

        <form onSubmit={handleSubmit}>
          <input
            type="text"
            name="tracking_id"
            placeholder="Tracking ID"
            value={formData.tracking_id}
            onChange={handleChange}
          />

          <input
            type="text"
            name="no_resi"
            placeholder="Nomor Resi"
            value={formData.no_resi}
            onChange={handleChange}
          />

          <input
            type="text"
            name="status"
            placeholder="Status Paket"
            value={formData.status}
            onChange={handleChange}
          />

          <input
            type="text"
            name="location"
            placeholder="Lokasi Saat Ini"
            value={formData.location}
            onChange={handleChange}
          />

          <textarea
            name="note"
            placeholder="Catatan"
            rows="4"
            value={formData.note}
            onChange={handleChange}
          />

          <button type="submit">
            Simpan Tracking
          </button>
        </form>
      </div>
    </div>
  );
}

export default Tracking;
