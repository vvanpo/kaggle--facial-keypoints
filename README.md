Kaggle - Facial Keypoints Recognition
========================
https://www.kaggle.com/c/facial-keypoints-detection

The supplied dataset includes 4 files:
 - test.csv
 - training.csv
 - IdLookupTable.csv
 - SampleSubmission.csv

test.csv is a list of image id numbers, following by the image specification in PGM format.
training.csv lists supplied keypoints to compare against the subsequent images also in PGM format.

To convert the PGM formats to something browsable, use the convert-pgm-png script.  The script takes either csv as input and creates PNG files in the present dir.

e.g.:

    $ # Ensure you have ImageMagick installed
    $ which convert
    /usr/bin/convert
    $ mkdir test-png && cd test-png
    $ ../convert-pgm-png < ../test.csv
    $ mkdir ../training-png && cd ../training-png
    $ ../convert-pgm-png < ../training.csv

You should now have two directories filled with PNG files.  PNG reduces dataset size by a factor of 3.
