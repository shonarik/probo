# Toward SOC 2

## Not everything is mandatory

You know better than anyone (including the auditor) what is best for your
company. You might have good reasons for performing or not a task as every
company is unique.

To help you evaluate the importance of each tasks you will run into, they are
labelled with three levels:

- **Mandatory** – The essential and fundamental elements. If you don’t have
  those, your auditor or customer will ask questions, you better justify it.
- **Optional** – Your auditor or customer might ask questions if any of those
  elements are necessary to mitigate a risk you have.

    <details>
    🗣
    
    A good example is penetration testing.
    
    - Penetration tests are expensive, but they can be a good investment, especially if you are running into a prospect requiring it
    - However they probably wont make sense if you are at the MVP stage and you are gonna trash and rebuild your product in a few month
    </details>

- **Advanced** – They show a great commitment toward security. Unless it is the
  only way to mitigate a risk very specific to your company, you won’t be asked
  about it.

## What’s next ?

We have regrouped what you need to do by different thematic in order to setup
the proper foundations for your company to get SOC-2:

<details>

**Pro tip:** Setup screenshot to clipboard You will have to take quite a lot of
screenshot until we automate most of it. We recommend that you setup a
screenshot to clipboard so you can just take a screenshot and paste-it saving
you ton of time. On mac:

- Use **CMD + SHIFT + 5** to enter screenshot mode
- Click options, to change “save to” to “clipboard”
- Now you can use **CMD + SHIFT + 3** to take a screenshot to clipboard
- And paste it in the right place with **CMD + V**

</details>

### Physical assets

Protect your physical environment to prevent data leaks or outages from
unauthorized access.

- [Secure your offices and internet access](physical/facilities/README.md)
- [Manage your computers](physical/hardware/README.md)

### Employees

Your team is your first line of defense—educate, empower, and secure them.

- [Set up your employees for success](personnel/lifecycle/README.md)
- [Secure your emails](personnel/comms/README.md)
- [Configure your system access](personnel/access/README.md)

### Core assets

These are the heart of your company—prioritize their security.

- [Secure your codebase](core/src/README.md)
- [Secure your infrastructure](core/infra/README.md)
- [Protect your network](core/network/README.md)
- [Safeguard your data](core/data/README.md)

### Alert & act

Be proactive and prepared—track activity and respond quickly to issues.

- [Log collection and monitoring](operations/monitoring/README.md)
- [Prepare for incidents](operations/incidents/README.md)

### Vendors

Keep your partnerships secure by managing third-party risks.

### Transparency

SOC 2 is about showing how you operate—document and share your processes.

- [Have a security page](TRA.001_have_a_security_page.md)
- [Clearly explain your services](transparency/TRA.002_clearly_explain_your_services.md)
- [External support available](TRA.003_external_support_available.md)

### Review and keep things up to date

#### Why does it matter?

Your company changes over time, and so should your security posture.

#### How can I proceed?

Some things don’t need review (eg: MFA is enabled), but your infrastructure and
your employee are changing.

⇒ You need to make sure people and digital asset with access are the one working
today in/with your company.

- [Remove unauthorized assets](core/infra/COR.INF.007_asset_decommissioning.md)
- [Conduct an access reviews](personnel/access/PER.ACC.006_periodic_access_review.md)
- [Test your disaster recovery plan](operations/incidents/OPS.INC.003_drp_testing.md)
