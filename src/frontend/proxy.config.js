var PROXY_CONFIG = [
  {
    context: [
      "/login",
      "/logout",
      "/currentuser",
      "/api",
      "/openapi",
      "/ws",
      "/healthz",
      "/session"
    ],
    // target: "http://192.168.112.120:17893",
    target: "http://192.168.112.26:3030",
    secure: false,
    changeOrigin: true
  }
];

module.exports = PROXY_CONFIG;
