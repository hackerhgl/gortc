## forgetPasswordSendOTP

- Check if an active user exist with the email. If it Does:
  - Check if user is verified.

- If user have not requested a code before then we create code & send e-mail.
- If user has requested a code before then following checks are in work:
  - Send email of code if it's requested after one minute and before expiry time of existing code.
  - Block request if user is requesting code again in one minute time frame.
  - Create a new code & send email if not active or the code has expired.