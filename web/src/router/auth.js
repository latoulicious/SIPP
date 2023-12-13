export default function (to, from, next) {
  const jwtToken = localStorage.getItem("jwtToken");
  console.log("Retrieved token:", jwtToken);

  // Check if the user is authenticated (you might want to add more checks)
  if (!jwtToken) {
    // Redirect to the login page if not authenticated
    next({ name: "login" }); // Use the name of the route
  } else {
    // Extract user role from the JWT token (modify this based on your token structure)
    const userRole = extractUserRole(jwtToken);

    // Check if the user has the required role to access the route
    if (to.name === "user" && userRole !== "admin") {
      // If the user is not an admin, redirect to a different route or show an error
      next({ name: "dashboard" }); // Redirect to the dashboard for example
    } else {
      // Proceed to the next route if authenticated and has the required role
      next();
    }
  }
}

// Modify this function based on the structure of your JWT token
function extractUserRole(jwtToken) {
  // Example: Decode the JWT token and extract the role
  const decodedToken = decodeToken(jwtToken);
  return decodedToken ? decodedToken.role : null;
}

function decodeToken(token) {
  // Modify this based on your JWT decoding mechanism (e.g., using a library)
  try {
    return JSON.parse(atob(token.split(".")[1]));
  } catch (error) {
    console.error("Error decoding JWT token:", error);
    return null;
  }
}
