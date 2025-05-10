import React, { useState, useEffect } from "react";
import "../style/Pagination.css"

function Pagination({currentPage, setCurrentPage, totalPages}) {
    return (
    <div className="pagination">
        <button
          className="page-btn"
          disabled={currentPage === 1}
          onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
        >
          Prev
        </button>

        {/* Always show first page */}
        <button
          className={`page-btn ${currentPage === 1 ? "active" : ""}`}
          onClick={() => setCurrentPage(1)}
        >
          1
        </button>

        {/* Ellipsis before current page if needed */}
        {currentPage > 3 && <span className="ellipsis">...</span>}

        {/* Current Page (only show if not 1 or last) */}
        {currentPage !== 1 && currentPage !== totalPages && (
          <button className="page-btn active" disabled>
            {currentPage}
          </button>
        )}

        {/* Ellipsis after current page if needed */}
        {currentPage < totalPages - 2 && <span className="ellipsis">...</span>}

        {/* Always show last page if not already shown */}
        {totalPages > 1 && (
          <button
            className={`page-btn ${currentPage === totalPages ? "active" : ""}`}
            onClick={() => setCurrentPage(totalPages)}
          >
            {totalPages}
          </button>
        )}

        <button
          className="page-btn"
          disabled={currentPage === totalPages}
          onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
        >
          Next
        </button>
    </div>
    )
}

export default Pagination;