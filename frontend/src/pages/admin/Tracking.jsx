import { useState } from "react";
import trackingService from "../../services/trackingService";
import Navbar from "../../components/Navbar";

// ───────────────── STATUS BADGE ─────────────────
const STATUS_MAP = {
  IN_TRANSIT: {
    bg: "#d97706",
    label: "Dalam Perjalanan",
  },
  OUT_FOR_DELIVERY: {
    bg: "#ea580c",
    label: "Dikirim Kurir",
  },
  DELIVERED: {
    bg: "#16a34a",
    label: "Terkirim",
  },
};

function Badge({ status }) {
  const s = STATUS_MAP[status] ?? {
    bg: "#6b7280",
    label: status,
  };

  return (
    <span
      style={{
        background: s.bg,
        color: "white",
        padding: "2px 10px",
        borderRadius: 20,
        fontSize: 11,
        fontWeight: 700,
      }}
    >
      {s.label}
    </span>
  );
}

// ───────────────── TIMELINE BUILDER ─────────────────
function buildTimeline(trackings) {
  return trackings.map((t, i) => ({
    icon:
      t.status === "DELIVERED"
        ? "✅"
        : t.status === "OUT_FOR_DELIVERY"
        ? "🚚"
        : "🚛",

    title:
      t.status === "DELIVERED"
        ? "Paket Diterima Penerima"
        : t.status === "OUT_FOR_DELIVERY"
        ? "Dikirim ke Penerima"
        : "Dalam Perjalanan",

    lokasi: t.location,

    waktu: t.created_at
      ? new Date(t.created_at).toLocaleString("id-ID")
      : "-",

    status: t.status,

    note: t.note,

    currentLocation: t.location,

    current: i === trackings.length - 1,

    done: true,
  }));
}

// ───────────────── COMPONENT ─────────────────
export default function Tracking() {

  const [noResi, setNoResi] = useState("");

  const [result, setResult] = useState([]);

  const [loading, setLoading] = useState(false);

  const [searched, setSearched] = useState("");

  const handleCari = async () => {

    const resi = noResi.trim();

    if (!resi) return;

    setLoading(true);

    setSearched(resi);

    try {

      const response =
        await trackingService.getTrackingByResi(
          resi
        );

      setResult(response.data || []);

    } catch (error) {

      console.error(error);

      setResult([]);

    }

    setLoading(false);
  };

  const timeline = buildTimeline(result);

  const currentStep =
    timeline.findLast?.((s) => s.current)
    ??
    timeline[timeline.length - 1];
  return (
    <>
      <Navbar />

      <div
        style={{
          padding: "24px",
          maxWidth: 920,
          margin: "0 auto",
        }}
      >
        <h1
          style={{
            marginBottom: 4,
            color: "#1a1a2e",
          }}
        >
          🚚 Tracking Paket
        </h1>

        <p
          style={{
            color: "#888",
            fontSize: 13,
            marginBottom: 24,
          }}
        >
          Lihat riwayat perjalanan paket berdasarkan nomor resi.
        </p>

        {/* PANEL TRACKING */}
        <div
          style={{
            background: "white",
            borderRadius: 14,
            padding: "20px 24px",
            boxShadow: "0 2px 10px rgba(0,0,0,0.07)",
            border: "1px solid #e5e7eb",
          }}
        >
          <h3
            style={{
              margin: "0 0 4px",
              fontSize: 15,
              color: "#1a1a2e",
            }}
          >
            🔍 Lacak Paket via No Resi
          </h3>

          <p
            style={{
              margin: "0 0 14px",
              fontSize: 12,
              color: "#888",
            }}
          >
            Masukkan nomor resi untuk melihat riwayat perjalanan paket.
          </p>

          <div
            style={{
              display: "flex",
              gap: 10,
            }}
          >
            <input
              type="text"
              value={noResi}
              placeholder="Contoh : TRK-20240001"
              onChange={(e) =>
                setNoResi(e.target.value)
              }
              onKeyDown={(e) =>
                e.key === "Enter" && handleCari()
              }
              style={{
                flex: 1,
                padding: "11px 14px",
                border: "1px solid #d1d5db",
                borderRadius: 8,
                fontSize: 14,
                outline: "none",
              }}
            />

            <button
              onClick={handleCari}
              disabled={loading}
              style={{
                padding: "11px 24px",
                background: loading
                  ? "#93c5fd"
                  : "#2563eb",
                color: "white",
                border: "none",
                borderRadius: 8,
                cursor: "pointer",
                fontWeight: 700,
              }}
            >
              {loading ? "Mencari..." : "Cari"}
            </button>
          </div>

          {searched !== "" && timeline.length === 0 && (
            <div
              style={{
                marginTop: 20,
                background: "#fef2f2",
                border: "1px solid #fecaca",
                padding: 16,
                borderRadius: 12,
              }}
            >
              <p
                style={{
                  color: "#dc2626",
                  margin: 0,
                  fontWeight: 700,
                }}
              >
                ❌ Paket tidak ditemukan
              </p>

              <p
                style={{
                  marginTop: 5,
                  fontSize: 12,
                  color: "#6b7280",
                }}
              >
                No resi {searched} belum memiliki data tracking.
              </p>
            </div>
          )}

          {currentStep && (
            <div
              style={{
                background:
                  "linear-gradient(135deg,#1d4ed8,#2563eb)",
                borderRadius: 12,
                padding: "14px 18px",
                color: "white",
                marginTop: 20,
                marginBottom: 20,
              }}
            >
              <p
                style={{
                  margin: 0,
                  fontSize: 12,
                  opacity: .8,
                }}
              >
                📍 POSISI PAKET SEKARANG
              </p>

              <p
                style={{
                  margin: "6px 0 0",
                  fontSize: 17,
                  fontWeight: 700,
                }}
              >
                {currentStep.currentLocation}
              </p>

              <p
                style={{
                  margin: "3px 0 0",
                  fontSize: 12,
                  opacity: .85,
                }}
              >
                {currentStep.title}
              </p>
            </div>
          )}
          {timeline.length > 0 && (
            <>
              <p
                style={{
                  fontWeight: 700,
                  fontSize: 13,
                  color: "#374151",
                  marginBottom: 12,
                }}
              >
                Riwayat Perjalanan Paket
              </p>

              {timeline.map((step, i) => (
                <div
                  key={i}
                  style={{
                    display: "flex",
                    gap: 14,
                  }}
                >
                  {/* ICON + GARIS */}
                  <div
                    style={{
                      display: "flex",
                      flexDirection: "column",
                      alignItems: "center",
                    }}
                  >
                    <div
                      style={{
                        width: 38,
                        height: 38,
                        borderRadius: "50%",
                        background: step.current
                          ? "#2563eb"
                          : "#eff6ff",

                        border: `2px solid ${
                          step.current
                            ? "#2563eb"
                            : "#93c5fd"
                        }`,

                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",

                        fontSize: 18,

                        boxShadow: step.current
                          ? "0 0 0 4px #bfdbfe"
                          : "none",
                      }}
                    >
                      {step.icon}
                    </div>

                    {i < timeline.length - 1 && (
                      <div
                        style={{
                          width: 2,
                          minHeight: 28,
                          background: "#bfdbfe",
                          marginTop: 5,
                        }}
                      />
                    )}
                  </div>

                  {/* ISI */}
                  <div
                    style={{
                      paddingTop: 5,
                      paddingBottom: 20,
                    }}
                  >
                    <div
                      style={{
                        display: "flex",
                        gap: 8,
                        alignItems: "center",
                        flexWrap: "wrap",
                      }}
                    >
                      <span
                        style={{
                          fontWeight: 700,
                          fontSize: 13,
                        }}
                      >
                        {step.title}
                      </span>

                      <Badge status={step.status} />

                      {step.current && (
                        <span
                          style={{
                            background: "#fef9c3",
                            color: "#854d0e",
                            border: "1px solid #fde047",
                            padding: "2px 8px",
                            borderRadius: 20,
                            fontSize: 10,
                            fontWeight: 700,
                          }}
                        >
                          ● SEKARANG
                        </span>
                      )}
                    </div>

                    <p
                      style={{
                        margin: "4px 0",
                        color: "#2563eb",
                        fontWeight: 600,
                        fontSize: 12,
                      }}
                    >
                      📍 {step.lokasi}
                    </p>

                    <p
                      style={{
                        margin: "2px 0",
                        color: "#6b7280",
                        fontSize: 11,
                      }}
                    >
                      📝 {step.note}
                    </p>

                    <p
                      style={{
                        margin: "2px 0",
                        color: "#6b7280",
                        fontSize: 11,
                      }}
                    >
                      🕒 {step.waktu}
                    </p>
                  </div>
                </div>
              ))}
            </>
          )}
        </div>
      </div>
    </>
  );
}