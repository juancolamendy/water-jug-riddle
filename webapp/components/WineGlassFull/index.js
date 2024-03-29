import * as React from "react";

function SvgComponent(props) {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 370 370" {...props}>
      <path d="M275.134 126.591L257.757 0H112.243L94.866 126.591c-5.252 24.172-3.581 48.57 4.721 68.771 9.483 23.08 26.773 39.796 49.998 48.343a102.047 102.047 0 0020.415 5.18V340h-35v30h100v-30h-35v-91.115a102.047 102.047 0 0020.415-5.18c23.225-8.547 40.515-25.263 49.998-48.343 8.302-20.201 9.973-44.599 4.721-68.771zM231.593 30l4.805 35H133.603l4.805-35h93.185z" />
    </svg>
  );
}

export default SvgComponent;

