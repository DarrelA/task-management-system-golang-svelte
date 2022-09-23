import { toast } from "@zerodevx/svelte-toast";

export const errorToast = (message) =>
  toast.push(message, {
    theme: {
      "--toastBackground": "#ba181b",
      "--toastColor": "white",
      "--toastBarBackground": "#242423",
      "--toastWidth": "300px",
      "--toastBorderRadius": "10px",
    },
    duration: 800,
  });

export const successToast = (message) =>
  toast.push(message, {
    theme: {
      "--toastBackground": "#52b788",
      "--toastColor": "black",
      "--toastBarBackground": "#242423",
      "--toastWidth": "300px",
      "--toastBorderRadius": "10px",
    },
    duration: 800,
  });
