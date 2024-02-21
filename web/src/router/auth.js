// auth.js

function requireAuth(to, from, next) {
  const jwtToken = localStorage.getItem("jwtToken");

  if (!jwtToken) {
    next({ name: "login" });
    return;
  }

  const decodedToken = decodeJwt(jwtToken);
  const isAdmin = decodedToken && decodedToken.role === "Admin";
  const isGuru = decodedToken && decodedToken.role === "Guru";
  const isWakasek =
    decodedToken && decodedToken.role === "Wakil Kepala Sekolah";
  const isKepsek = decodedToken && decodedToken.role === "Kepala Sekolah";

  if (isAdmin || isGuru || isWakasek || isKepsek) {
    next();
  } else {
    next("/unauthorized");
  }
}

function decodeJwt(jwtToken) {
  try {
    const [, payloadBase64] = jwtToken.split(".");
    const payload = JSON.parse(atob(payloadBase64));

    return payload;
  } catch (error) {
    console.error("Error decoding JWT:", error);
    return null;
  }
}

export default requireAuth;
