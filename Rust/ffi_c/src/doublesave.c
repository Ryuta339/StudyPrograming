#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "mat.h"

int doublesave (const char * filename, const double * arry, size_t row, size_t col) {
	MATFile *pmat;
	mxArray * pa;
	int status;

	printf ("Saving matrix...\n");
	for (size_t c=0; c<col; c++) {
		for (size_t r=0; r<row; r++)
			printf ("\t%2.5lf", arry[c*row+r]);
		printf ("\n");
	}

	printf ("Creating file %s...\n\n", filename);
	if ((pmat = matOpen (filename, "w")) == NULL) {
		fprintf (stderr, "Error creating file %s.\n", filename);
		return (EXIT_FAILURE);
	}

	pa = mxCreateDoubleMatrix (row, col, mxREAL);
	if (pa == NULL) {
		fprintf (stderr, "%s: Out pf memory on line %d.\n", __FILE__, __LINE__);
		fprintf (stderr, "Unable to create mxArray.\n");
		return (EXIT_FAILURE);
	}

	double * ptr = (double *)(mxGetDoubles(pa));
	for (size_t c=0; c<col; c++) {
		for (size_t r=0; r<row; r++)
			ptr[c*row+r] = arry[c*row+r];
	}

	status = matPutVariable (pmat, "arry", pa);
	if (status != 0) {
		fprintf (stderr, "%s: Error using matPutVariable on line %d.\n", __FILE__, __LINE__);
		return (EXIT_FAILURE);
	}

	/* clean up */
	mxDestroyArray (pa);
	if (matClose (pmat) != 0) {
		fprintf (stderr, "Error closing file %s.\n", filename);
		return (EXIT_FAILURE);
	}

	printf ("Done.\n");
	return EXIT_SUCCESS;
}
