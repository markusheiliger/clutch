import React from "react";
import { useLocation } from "react-router-dom";
import type { clutch as IClutch } from "@clutch-sh/api";

import { useNavigate } from "../navigation";
import { client } from "../Network";
import type { ClutchError } from "../Network/errors";

interface ShortLinkHydratorProps {
  hydrate: (HydrateData) => void;
}

/**
 * Component that will be present for a route which will look for a short link
 * - If found
 * - It will call down to the API with the hash and ask for any data pertaining to it
 * - If the API call is successful
 *   - It will use the given hydrate function to send the returned state off to the StorageContext
 *   - It will navigate to the route given in the returned state
 * - If the API call is not successful
 *   - It will leave a warning message in the console
 *   - Then navigate back to the home page
 */
const ShortLinkHydrator = ({ hydrate }: ShortLinkHydratorProps) => {
  const { pathname } = useLocation();
  const navigate = useNavigate();

  React.useEffect(() => {
    // Looking for a route similar to: "/sl/1234"
    const matches = pathname.match(/(\/sl\/)(.*)/i);
    if (matches && matches[2]) {
      const requestData: IClutch.shortlink.v1.IGetRequest = { hash: matches[2] };

      client
        .post("/v1/shortlink/get", requestData)
        .then(response => {
          const { path = "/", state } = response.data as IClutch.shortlink.v1.IGetResponse;
          hydrate(state);
          navigate(path);
        })
        .catch((error: ClutchError) => {
          // eslint-disable-next-line no-console
          console.warn(`Shortlink ${matches[2]} errored, redirecting home`, error?.message);
          navigate("/");
        });
    }
  }, [pathname]);

  // currently return null so that nothing is rendered
  // TODO: either make a loading spinner or something to go here if the API calls are not responsive enough
  return null;
};

export default ShortLinkHydrator;
